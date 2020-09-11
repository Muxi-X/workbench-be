package service

import (
	"context"

	"muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

// Remove … 移出团队内成员
func (ts *TeamService) Remove(ctx context.Context, req *pb.RemoveRequest, res *pb.Response) error {
	if err := RemoveFromTeam(req.TeamId, req.UserId); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

// RemoveFromTeam remove from team
func RemoveFromTeam(teamID uint32, userID uint32) error {
	users := []uint32{userID}
	if err := UpdateUsersGroupIDOrTeamID(users, 0, model.TEAM); err != nil {
		return err
	}
	if err := model.TeamCountOperation(teamID, 1, model.SUB); err != nil {
		return err
	}
	return nil
}
