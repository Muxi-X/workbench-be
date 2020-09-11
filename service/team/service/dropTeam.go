package service

import (
	"context"

	"muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

// DropTeam … 删除团队
func (ts *TeamService) DropTeam(ctx context.Context, req *pb.DropTeamRequest, res *pb.Response) error {
	// 获取usersid
	usersID, err := GetUsersIdByTeamId(req.TeamId)
	if err != nil {
		return e.ServerErr(errno.ErrClient, err.Error())
	}

	// 将对应相应team内的所有成员， teamid置零
	if err := UpdateUsersGroupIDOrTeamID(usersID, model.NOTEAM, model.TEAM); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 删除团队
	err = model.DropTeam(req.TeamId)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
