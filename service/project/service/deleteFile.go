package service

import (
	"context"
	// apb "muxi-workbench-attention/proto"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
	"time"
)

// DeleteFile ... 删除文件
func (s *Service) DeleteFile(ctx context.Context, req *pb.DeleteRequest, res *pb.Response) error {
	item, err := model.GetFile(req.Id)
	if err != nil {
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
	if err = model.DeleteFile(m.DB.Self, trashbin, fatherId, isFatherProject); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 向取消关注发起请求
	err = DeleteAttentionsFromAttentionService(req.Id, uint32(constvar.FileCode))
	if err != nil {
		return err
	}

	return nil
}
