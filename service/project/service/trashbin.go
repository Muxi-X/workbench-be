package service

import (
	"fmt"
	"muxi-workbench-project/model"
	"muxi-workbench/log"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// GoTidyTrashbin ... 子协程定时任务
// 清理用户删除文件 和 自己过期的文件
// 1. 查找 trashbin 表 re 为 1 记录 和 expire_at 小于 now 记录
// 2. 删除回收站记录、修改原表 re 字段、同步删除 redis
func GoTidyTrashbin(db *gorm.DB) {
	tidyDay := viper.GetInt("trashbin.tidy_time")
	if tidyDay == 0 {
		log.Error("tidyDay failed to get")
		return
	}

	tidyDuration := time.Hour * time.Duration(tidyDay)

	for {
		if err := model.TidyTrashbin(db); err != nil {
			log.Error("TidyTrashbin error", zap.String("cause", err.Error()))
		}

		time.Sleep(tidyDuration)
	}
}

// SynchronizeTrashbinToRedis ... 开启服务调用同步 redis
func SynchronizeTrashbinToRedis() error {
	list, err := model.GetAllTrashbin()

	var res []string
	for _, v := range list {
		// 修改原表 re 字段 和 获取子文件
		switch v.FileType {
		case constvar.DocCode:
			res = append(res, fmt.Sprintf("%d-%d", v.FileId, constvar.DocCode))
		case constvar.FileCode:
			res = append(res, fmt.Sprintf("%d-%d", v.FileId, constvar.FileCode))
		case constvar.DocFolderCode:
			err = model.GetDocChildFolder(v.FileId, &res)
		case constvar.FileFolderCode:
			err = model.GetFileChildFolder(v.FileId, &res)
		}

		if err != nil {
			return err

		}
	}

	// 同步 redis
	if len(res) != 0 {
		if err := m.SAddToRedis(constvar.Trashbin, res); err != nil {
			return err
		}
	}

	return nil
}
