package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
	"time"

	"github.com/jinzhu/gorm"
)

// DeleteDoc ... 删除文档
// 插入回收站 同步 redis
func (s *Service) DeleteDoc(ctx context.Context, req *pb.DeleteRequest, res *pb.ProjectIDResponse) error {
	// 找 name 和 creatorId
	item, err := model.GetDoc(req.Id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return e.ServerErr(errno.ErrNotFound, err.Error())
		}
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 权限判定
	if item.CreatorID != req.UserId {
		if req.Role <= constvar.Normal {
			return e.BadRequestErr(errno.ErrPermissionDenied, "")
		}
	}

	// 获取 fatherId
	isFatherProject := false
	var fatherId uint32
	if item.FatherId == 0 { // fatherId 为 0 则是 project
		isFatherProject = true
		fatherId = item.ProjectID
	} else {
		fatherId = item.FatherId
	}

	trashbin := &model.TrashbinModel{
		FileId:     req.Id,
		FileType:   constvar.DocCode,
		Name:       item.Name,
		DeleteTime: time.Now().Format("2006-01-02 15:04:05"),
		CreateTime: item.CreateTime,
	}

	// 事务
	err = model.DeleteDoc(m.DB.Self, trashbin, fatherId, isFatherProject)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}
