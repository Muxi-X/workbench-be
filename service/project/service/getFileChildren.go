package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// GetFileChildren ... 获取任意文件夹目录下的文件树
func (s *Service) GetFileChildren(ctx context.Context, req *pb.GetRequest, res *pb.Children) error {
	// 新增判断节点是否被删
	// 文件夹，只需要查自己有无被删
	isDeleted, err := model.AdjustSelfIfExist(req.Id, constvar.FileFolderCode)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if isDeleted { // 存在 redis 返回 1, 说明被删
		return e.NotFoundErr(errno.ErrNotFound, "This file has been deleted.")
	}

	item, err := model.GetFileChildrenById(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Children = item.Children

	return nil
}
