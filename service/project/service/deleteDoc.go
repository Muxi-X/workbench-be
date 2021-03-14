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
// 用接口完成文档删除和文件树修改
func (s *Service) DeleteDoc(ctx context.Context, req *pb.DeleteRequest, res *pb.ProjectIDResponse) error {
	// 先查找再删除
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

	item.Re = true
	res.Id = item.ProjectID

	// 软删除,修改 re 字段
	// 事务
	err = model.DeleteDoc(m.DB.Self, item, req.FatherId, req.ChildrenPositionIndex, req.FatherType)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}
