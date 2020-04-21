package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// UpdateProjectInfo ... 更新项目信息
func (s *Service) UpdateProjectInfo(ctx context.Context, req *pb.ProjectInfo, res *pb.Response) error {

	item, err := model.GetProject(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	item.Name = req.Name
	item.Intro = req.Intro

	if err := item.Update(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}
