package service

import (
	"context"

	ppb "muxi-workbench-project/proto"
	upb "muxi-workbench-user/proto"

	"github.com/micro/go-micro"
)

// FeedService ... 动态服务
type FeedService struct{}

// GetFilterFromProjectService get filter data from project-service
func GetFilterFromProjectService(id uint32) ([]uint32, error) {
	service := micro.NewService()
	service.Init()

	client := ppb.NewProjectServiceClient("workbench.service.project", service.Client())

	rsp, err := client.GetProjectIdsForUser(context.Background(), &ppb.GetRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return rsp.List, nil
}

// GetInfoFromUserService get user's name and avatar from user-service
func GetInfoFromUserService(id uint32) (string, string, error) {
	service := micro.NewService()
	service.Init()

	client := upb.NewUserServiceClient("workbench.service.user", service.Client())

	rsp, err := client.GetProfile(context.Background(), &upb.GetRequest{Id: id})
	if err != nil {
		return "", "", err
	}

	return rsp.Name, rsp.AvatarUrl, nil
}
