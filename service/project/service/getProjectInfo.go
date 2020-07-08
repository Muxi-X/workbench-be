package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
)

// GetProjectInfo ... 获取项目信息
func (s *Service) GetProjectInfo(ctx context.Context, req *pb.GetRequest, res *pb.ProjectInfo) error {

	project, err := model.GetProject(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	count, err := model.GetProjectUserCount(req.Id)

	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = project.ID
	res.Name = project.Name
	res.Intro = project.Intro
	res.UserCount = count

	return nil
}
