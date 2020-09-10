package service

import (
	"context"
	"fmt"
	"muxi-workbench-team/model"
	upb "muxi-workbench-user/proto"
)

// TeamService … 团队服务
type TeamService struct {
}

// Init other service
func Init() {
	UserInit(UserService, UserClient)
}

// UpdateUsersGroupIDOrTeamID update user's group_id by user_id
func UpdateUsersGroupIDOrTeamID(usersID []uint32, value uint32, kind uint32) error {
	fmt.Println(UserClient)
	_, err := UserClient.UpdateTeamAndGroupForUsers(context.Background(), &upb.UpdateTeamGroupRequest{Ids: usersID, Value: value, Kind: kind})

	if err != nil {
		return err
	}

	return nil
}

// GetUsersIdByGroupid get users' infos by groupid
func GetUsersIdByGroupid(groupID uint32) ([]uint32, error) {
	rsp, err := UserClient.List(context.Background(), &upb.ListRequest{
		LastId: 0,
		Offset: 0,
		Limit:  0,
		Team:   model.MUXI,
		Group:  groupID,
	})
	if err != nil {
		return nil, err
	}

	users := make([]uint32, 0)
	for _, item := range rsp.List {
		users = append(users, item.Id)
	}
	return users, nil
}

// GetUsersIdByTeamId get users' infos by teamid
func GetUsersIdByTeamId(teamID uint32) ([]uint32, error) {
	rsp, err := UserClient.List(context.Background(), &upb.ListRequest{
		LastId: 0,
		Offset: 0,
		Limit:  0,
		Team:   teamID,
	})
	if err != nil {
		return nil, err
	}

	users := make([]uint32, 0)
	for _, item := range rsp.List {
		users = append(users, item.Id)
	}
	return users, nil
}

// GetMemberInfo list all members of a group
func GetMemberInfo(groupID uint32, limit uint32, offset uint32, pagination bool) ([]*MemberModel, uint64, error) {
	if !pagination {
		offset = 0
	}

	memberlist := make([]*MemberModel, 0)

	rsp, err := UserClient.List(context.Background(), &upb.ListRequest{
		LastId: 0,
		Offset: offset,
		Limit:  limit,
		Team:   model.MUXI,
		Group:  groupID,
	})
	if err != nil {
		return memberlist, 0, err
	}
	count := uint64(rsp.Count)

	group, err := model.GetGroup(groupID)
	if err != nil {
		return memberlist, count, err
	}

	for _, item := range rsp.List {
		memberlist = append(memberlist, &MemberModel{
			UserID:    item.Id,
			TeamID:    item.Team,
			GroupID:   item.Group,
			GroupName: group.Name,
			Role:      item.Role,
			Email:     item.Email,
			Avatar:    item.Avatar,
			Name:      item.Name,
		})
	}
	return memberlist, count, nil
}

// GetUsersByApplys get users by applys ID
func GetUsersByApplys(applys []*model.ApplyModel, count uint64) ([]*model.ApplyUserItem, uint64, error) {
	userIDs := GetUsersIDByApplys(applys)

	applyuserList := make([]*model.ApplyUserItem, 0)

	rsp, err := UserClient.GetInfo(context.Background(), &upb.GetInfoRequest{Ids: userIDs})
	if err != nil {
		return applyuserList, 0, err
	}

	for _, item := range rsp.List {
		applyuserList = append(applyuserList, &model.ApplyUserItem{
			Name:  item.Name,
			ID:    item.Id,
			Eamil: "",
		})
	}

	return applyuserList, count, nil
}
