package service

import (
	"context"
	e "github.com/Muxi-X/workbench-be/pkg/err"
	"github.com/Muxi-X/workbench-be/service/team/errno"
	"github.com/Muxi-X/workbench-be/service/team/model"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
)

func (ts *TeamService) Join(ctx context.Context, req *pb.JoinRequest, res *pb.Response) error {
	//权限判断
	if req.Role != model.SUPERADMIN {

	}


	if err := model.JoinTeam(req.TeamId,req.UserId); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil

}