package service

import (
	"context"

	ppb "muxi-workbench-project/proto"
	upb "muxi-workbench-user/proto"

	"github.com/micro/go-micro"
)

var UserService micro.Service
var UserClient upb.UserServiceClient
var ProjectService micro.Service
var ProjectClient ppb.ProjectServiceClient

func UserInit() {
	UserService = micro.NewService(micro.Name("workbench.cli.user"))

	UserService.Init()

	UserClient = upb.NewUserServiceClient("workbench.service.user", UserService.Client())
}

func ProjectInit() {
	ProjectService = micro.NewService(micro.Name("workebnch.cli.project"))

	ProjectService.Init()

	ProjectClient = ppb.NewProjectServiceClient("workbench.service.project", ProjectService.Client())
}

// FeedService ... 动态服务
type FeedService struct{}

// GetFilterFromProjectService get filter data from project-service
func GetFilterFromProjectService(id uint32) ([]uint32, error) {
	rsp, err := ProjectClient.GetProjectIdsForUser(context.Background(), &ppb.GetRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return rsp.List, nil
}

// GetInfoFromUserService get user's name and avatar from user-service
func GetInfoFromUserService(id uint32) (string, string, error) {
	rsp, err := UserClient.GetProfile(context.Background(), &upb.GetRequest{Id: id})
	if err != nil {
		return "", "", err
	}

	return rsp.Name, rsp.Avatar, nil
}
