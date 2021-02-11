package service

import (
	"context"
	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	e "muxi-workbench/pkg/err"
	"time"
)

// CreateProject ... 建立项目
func (s *Service) CreateProject(ctx context.Context, req *pb.CreateProjectRequest, res *pb.ProjectIDResponse) error {
	t := time.Now()

	project := &model.ProjectModel{
		Name:   req.Name,
		Intro:  req.Intro,
		TeamID: req.TeamId,
		Time:   t.Format("2006-01-02 15:04:05"),
	}

	if err := project.Create(); err != nil {
		e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Id = project.ID

	return nil
}
