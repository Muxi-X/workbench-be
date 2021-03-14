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

// DeleteFile ... 删除文件 只修改 re 字段
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

	item.Re = true

	// 软删除，修改 re 字段
	// 事务
	err = model.DeleteFile(m.DB.Self, item, req.FatherId, req.ChildrenPositionIndex, req.FatherType)
	if err := item.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}
