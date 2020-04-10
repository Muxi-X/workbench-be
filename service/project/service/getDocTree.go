package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetDocTree ... 获取项目的文档树
func (s *Service) GetDocTree(ctx context.Context, req *pb.GetRequest, res *pb.Tree) error {

	project, err := model.GetProject(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Tree = project.DocTree

	return nil
}
