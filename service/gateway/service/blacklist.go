package service

import (
	"time"

	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/model"
	m "muxi-workbench/model"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// CheckInBlacklist checks whether a token is in blacklist
func CheckInBlacklist(token string) (bool, error) {
	var exist bool
	var err error

	// 先查看 redis 中是否有该数据，若有，则在黑名单中
	exist, err = m.HasExistedInRedis(token)
	if err != nil {
		log.Error("HasExistedInRedis error", zap.String("cause", err.Error()))
		return false, err
	} else if exist {
		return true, nil
	}

	// 查看 MySQL 中该数据是否在黑名单中
	exist, err = model.HasExistedInBlacklist(token)
	if err != nil {
		log.Error("HasExistedInBlacklist error", zap.String("cause", err.Error()))
		return false, err
	}
	return exist, nil
}

// AddToBlacklist adds a token into blacklist
func AddToBlacklist(token string, expiresAt int64) error {
	// 加入 MySQL 黑名单

	record := &model.BlacklistModel{
		Token:     token,
		ExpiresAt: expiresAt,
	}
	if err := record.Create(); err != nil {
		return err
	}

	// 加入 redis 黑名单中

	// 过期时间(s)
	var expiration time.Duration = time.Duration((expiresAt - time.Now().Unix())) * time.Second

	if err := m.SetStringInRedis(token, 1, expiration); err != nil {
		return err
	}

	return nil
}

// TidyBlacklist ... 定时清理过期的记录
func TidyBlacklist() {
	tidyDay := viper.GetInt("blacklist.tidy_time")
	if tidyDay == 0 {
		log.Error("tidyDay failed to get")
		return
	}

	tidyDuration := time.Hour * time.Duration(tidyDay)

	for {
		if err := model.DeleteExpiredBlacklist(); err != nil {
			log.Error("TidyBlacklist error", zap.String("cause", err.Error()))
		}

		time.Sleep(tidyDuration)
	}
}
