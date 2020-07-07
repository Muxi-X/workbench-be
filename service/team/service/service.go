package service

import (
	"github.com/micro/go-micro"
	upb "github.com/Muxi-X/workbench-be/service/user/proto"
)

//TeamService … 团队服务
type TeamService struct {
}

//UpdateUsersInfo update user's group_id by user_id
func UpdateUsersGroupid(userid uint32,groupid uint32) error {
	service := micro.NewService()
	service.Init()

	client := upb.NewUserServiceClient("workbench.service.user", service.Client())

	//rsp, err := client.
	return nil
}

//GetUserInfo get users' infos by groupid
func GetUserInfo(id uint32) error {
	service := micro.NewService()
	service.Init()

	client := upb.NewUserServiceClient("workbench.service.user", service.Client())

	rsp, err := client.GetInfo()
	rsp.
}

//
func UpdateUserTeamId(userid uint32,teamid uint32) error {
	service := micro.NewService()
	service.Init()

	client := upb.NewUserServiceClient("workbench.service.user", service.Client())

	rsp, err := client.

}





