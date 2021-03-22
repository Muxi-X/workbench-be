package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// DeleteDoc ... 删除文档
// 插入回收站 同步 redis
func (s *Service) DeleteDoc(ctx context.Context, req *pb.DeleteRequest, res *pb.ProjectIDResponse) error {
	// 找 name 和 creatorId
	item, err := model.GetDoc(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 权限判定
	if item.CreatorID != req.UserId {
		if req.Role <= constvar.Normal {
			return e.BadRequestErr(errno.ErrPermissionDenied, "")
		}
	}

	trashbin := &model.TrashbinModel{
		FileId:   req.Id,
		FileType: constvar.DocCode,
		Name:     item.Name,
	}

	// 事务
	err = model.DeleteDoc(m.DB.Self, trashbin, req.FatherId, req.ChildrenPositionIndex, req.FatherType)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}
