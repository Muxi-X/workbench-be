package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetFileTree ... 获取项目的文件树
func (s *Service) GetFileTree(ctx context.Context, req *pb.GetRequest, res *pb.Tree) error {

	project, err := model.GetProject(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Tree = project.FileTree

	return nil
}
