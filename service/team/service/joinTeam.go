package service

import (
	"context"
	"muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

// Join …… 加入团队
func (ts *TeamService) Join(ctx context.Context, req *pb.JoinRequest, res *pb.Response) error {
	if err := JoinTeam(req.TeamId, req.UserId); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil

}

// JoinTeam join team
func JoinTeam(teamID uint32, userID uint32) error {
	users := []uint32{userID}
	if err := UpdateUsersGroupIDOrTeamID(users, teamID, model.TEAM); err != nil {
		return err
	}
	if err := model.TeamCountOperation(teamID, 1, model.ADD); err != nil {
		return err
	}
	return nil
}
