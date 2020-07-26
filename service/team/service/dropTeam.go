package service

import (
	"context"
	errno "muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

//Drop … Team 删除团队
func (ts *TeamService) DropTeam(ctx context.Context, req *pb.DropTeamRequest, res *pb.Response) error {
	//判断权限
	if req.Role != model.SUPERADMIN {
		return e.ServerErr(errno.ErrPermissionDenied, "权限不够")
	}

	//获取usersid
	usersid, err := model.GetUsersIdByTeamId(req.TeamId)
	if err := model.UpdateUsersGroupidOrTeamid(usersid, model.NOTEAM, model.TEAM); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	//删除组别
	err = model.DropTeam(req.TeamId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}