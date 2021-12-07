package service

import (
	"context"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"

	"muxi-workbench/pkg/handler"

	upb "muxi-workbench-user/proto"

	"github.com/micro/go-micro"
)

var UserService micro.Service
var UserClient upb.UserServiceClient

func UserInit() {
	UserService = micro.NewService(micro.Name("workbench.cli.user"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()),
	)

	UserService.Init()

	UserClient = upb.NewUserServiceClient("workbench.service.user", UserService.Client())
}
// StatusService ... 动态服务
type StatusService struct {
}

// GetInfoFromUserService get user's name and avatar from user-service
func GetInfoFromUserService(id uint32) (string, string, error) {
	rsp, err := UserClient.GetProfile(context.Background(), &upb.GetRequest{Id: id})
	if err != nil {
		return "", "", err
	}

	return rsp.Name, rsp.Avatar, nil
}
