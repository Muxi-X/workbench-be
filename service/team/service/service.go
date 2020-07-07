package service

import (
	"context"
	"github.com/Muxi-X/workbench-be/service/team/model"
	upb "github.com/Muxi-X/workbench-be/service/user/proto"
	"github.com/micro/go-micro"
)

//TeamService … 团队服务
type TeamService struct {
}

//UpdateUsersInfo update user's group_id by user_id
func UpdateUsersGroupidOrTeamid(userid []uint32,value uint32, kind uint32) error {
	service := micro.NewService()
	service.Init()

	client := upb.NewUserServiceClient("workbench.service.user", service.Client())

	_, err := client.UpdateTeamAndGroupForUsers(context.Background(), &upb.UpdateTeamGroupRequest{Ids:userid, Value:value, Kind:kind})
    if err != nil {
    	return err
	}

	return nil
}

//GetUserInfo get users' infos by groupid
func GetUsersId(groupid uint32) ([]uint32, error) {
	service := micro.NewService()
	service.Init()

	client := upb.NewUserServiceClient("workbench.service.user", service.Client())

	rsp, err := client.List(context.Background(),&upb.ListRequest{
		LastId:               0,
		Offset:               0,
		Limit:                0,
		Team:                 model.MUXI,
		Group:                groupid,
	})
	if err != nil {
		return nil, err
	}

	users := make([]uint32, 0)
	for _, item := range rsp.List {
		users = append(users, item.Id)
	}
	return users,nil
}

//list all members of a group
func GetMemberInfo(groupid uint32,limit uint32, offset uint32,pagination bool) ([]*model.MemberModel, uint64, error) {
	service := micro.NewService()
	service.Init()

	client := upb.NewUserServiceClient("workbench.service.user", service.Client())

    if !pagination {
    	limit = 0
    	offset = 0
	}

	memberlist := make([]*model.MemberModel, 0)

	rsp, err := client.List(context.Background(),&upb.ListRequest{
		LastId:               0,
		Offset:               limit,
		Limit:                offset,
		Team:                 model.MUXI,
		Group:                groupid,
	})
	count := uint64(rsp.Count)
	if err != nil {
		return memberlist, count, err
	}

	for _, item := range rsp.List {
		memberlist = append(memberlist, &model.MemberModel{
			UserID:    item.Id,
			TeamID:    item.Team,
			GroupID:   item.Group,
			GroupName: "",
			Role:      item.Role,
			Email:     "",
			Avatar:    item.Avatar,
			Name:      item.Name,
		})
	}
    return memberlist, count, nil

}






