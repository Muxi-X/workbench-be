package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
	"time"
)

// CreateProject ... 创建项目
func (s *Service) CreateProject(ctx context.Context, req *pb.CreateProjectRequest, res *pb.GetRequest) error {
	t := time.Now()

	project := model.ProjectModel{
		Name:   req.Name,
		TeamID: req.TeamId,
		Intro:  req.Desc,
		Time:   t.Format("2006-01-02 15:04:05"),
	}

	if err := project.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}
