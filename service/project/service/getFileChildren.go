package service

import (
	"context"
	"fmt"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	"muxi-workbench/pkg/constvar"
	e "muxi-workbench/pkg/err"
)

// GetFileChildren ... 获取任意文件夹目录下的文件树
func (s *Service) GetFileChildren(ctx context.Context, req *pb.GetRequest, res *pb.Children) error {
	// 新增判断节点是否被删
	// 文件夹，只需要查自己有无被删
	target := fmt.Sprintf("%d-%d", req.Id, constvar.FileFolderCode)
	isDeleted, err := m.SIsmembersFromRedis(constvar.Trashbin, target)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	if isDeleted {
		return e.BadRequestErr(errno.ErrDatabase, "This file has been deleted.")
	}

	item, err := model.GetFileChildrenById(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Children = item.Children

	return nil
}
