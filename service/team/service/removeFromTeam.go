package service

import (
	"context"
	e "github.com/Muxi-X/workbench-be/pkg/err"
	"github.com/Muxi-X/workbench-be/service/team/errno"
	"github.com/Muxi-X/workbench-be/service/team/model"
	pb "github.com/Muxi-X/workbench-be/service/team/proto"
)

func (ts *TeamService) Remove(ctx context.Context, req *pb.RemoveRequest, res *pb.Response) error {
	//权限判断
	if req.Role != model.SUPERADMIN {
		return e.ServerErr(errno.ErrPermissionDenied, "权限不够")
	}

	if err := model.RemoveformTeam(req.TeamId,req.UserId); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
