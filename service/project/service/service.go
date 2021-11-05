package service

import (
	"context"
	upb "muxi-workbench-user/proto"
)

// Service ... 项目服务
type Service struct {
}

var UserClient upb.UserServiceClient

// GetInfoFromUserService get user's name and avatar from user-service
func GetInfoFromUserService(id uint32) (string, error) {
	rsp, err := UserClient.GetProfile(context.Background(), &upb.GetRequest{Id: id})
	if err != nil {
		return "", err
	}

	return rsp.Name, nil
}
