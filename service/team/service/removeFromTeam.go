package service

import (
	"context"
	"errors"
	"muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	upb "muxi-workbench-user/proto"
	e "muxi-workbench/pkg/err"
)

// Remove … 移出团队内成员
func (ts *TeamService) Remove(ctx context.Context, req *pb.RemoveRequest, res *pb.Response) error {
	if err := RemoveFromTeam(req.TeamId, req.UserList); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

// RemoveFromTeam remove from team
func RemoveFromTeam(teamID uint32, userIDs []uint32) error {
	for _, id := range userIDs {
		info, err := UserClient.GetProfile(context.Background(), &upb.GetRequest{Id: id})
		if err != nil {
			return err
		}
		if info.Team != teamID {
			return errors.New("被移除成员中存在teamID与管理员teamID不符的情况")
		}
	}

	if err := UpdateUsersGroupIdOrTeamId(userIDs, 0, model.TEAM); err != nil {
		return err
	}
	if err := model.TeamCountOperation(teamID, uint32(len(userIDs)), model.SUB); err != nil {
		return err
	}
	return nil
}
