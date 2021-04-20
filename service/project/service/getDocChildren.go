package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// GetDocChildren ... 获取任意文档夹目录下的文档树
func (s *Service) GetDocChildren(ctx context.Context, req *pb.GetRequest, res *pb.Children) error {
	// 新增判断节点是否被删
	// 文件夹，只需要查自己有无被删
	isDeleted, err := model.AdjustSelfIfExist(req.Id, constvar.DocFolderCode)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if isDeleted {
		return e.NotFoundErr(errno.ErrNotFound, "This file has been deleted.")
	}

	item, err := model.GetDocChildrenById(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Children = item.Children

	return nil
}
