package service

import (
	"context"
	"errors"
	errno "muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

//Join …… 加入团队
func (ts *TeamService) Join(ctx context.Context, req *pb.JoinRequest, res *pb.Response) error {
	//权限判断
	if req.Role != model.SUPERADMIN && req.Role != model.ADMIN {
		err := errors.New("permission denied")
		return err
	}

	if err := model.JoinTeam(req.TeamId, req.UserId); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil

}
