package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// UpdateFileTree ... 更新项目的文件树
func (s *Service) UpdateFileTree(ctx context.Context, req *pb.UpdateTreeRequest, res *pb.Response) error {

	project, err := model.GetProject(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	project.FileTree = req.Tree

	if err := project.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}
