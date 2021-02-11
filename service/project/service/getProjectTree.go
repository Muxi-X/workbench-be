package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetProjectTree ... 获取任意项目下的两树
func (s *Service) GetProjectTree(ctx context.Context, req *pb.GetRequest, res *pb.ProjectTree) error {

	item, err := model.GetProjectChildrenById(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.DocTree = item.DocChildren
	res.FileTree = item.FileChildren

	return nil
}
