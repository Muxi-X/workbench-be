package service

import (
	"context"
	"time"

	"muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

// CreateTeam … 创建团队
func (ts *TeamService) CreateTeam(ctx context.Context, req *pb.CreateTeamRequest, res *pb.Response) error {
	if req.Role != model.SUPERADMIN && req.Role != model.ADMIN {
		return e.ServerErr(errno.ErrPermissionDenied, "权限不够")
	}

	t := time.Now()
	team := &model.TeamModel{
		Name:      req.TeamName,
		CreatorID: req.CreatorId,
		Time:      t.Format("2006-01-02 15:04:05"),
		Count:     1,
	}
	if err := team.Create(); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}
