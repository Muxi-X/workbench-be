package service

import (
	"context"

	"muxi-workbench-team/errno"
	"muxi-workbench-team/model"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

// GetApplications …… 列举申请
func (ts *TeamService) GetApplications(ctx context.Context, req *pb.ApplicationListRequest, res *pb.ApplicationListResponse) error {
	teamID, err := GetTeamIdByUserId(req.UserId)
	if err != nil {
		return e.ServerErr(errno.ErrClient, err.Error())
	}

	applys, count, err := model.ListApplys(req.Offset, req.Limit, req.Pagination, teamID)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	userList, _, err := GetUsersByApplys(applys, count)
	if err != nil {
		return e.ServerErr(errno.ErrClient, err.Error())
	}

	resList := make([]*pb.ApplyUserItem, 0)

	for _, item := range userList {
		resList = append(resList, &pb.ApplyUserItem{
			Id:    item.ID,
			Name:  item.Name,
			Email: item.Email,
		})
	}

	res.Count = uint32(count)
	res.List = resList

	return nil
}

// GetUsersIDByApplys get usersID from applys
func GetUsersIDByApplys(applys []*model.ApplyModel) []uint32 {
	ret := make([]uint32, 0)
	for _, value := range applys {
		ret = append(ret, value.UserID)
	}
	return ret
}
