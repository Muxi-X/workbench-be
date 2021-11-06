package service

import (
	"context"
	"github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"muxi-workbench/pkg/handler"

	upb "muxi-workbench-user/proto"
)

// Service ... 项目服务
type Service struct {
}

var UserClient upb.UserServiceClient
var UserService micro.Service

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

// GetInfoFromUserService get user's name and avatar from user-service
func GetInfoFromUserService(id uint32) (string, error) {
	rsp, err := UserClient.GetProfile(context.Background(), &upb.GetRequest{Id: id})
	if err != nil {
		return "", err
	}

	return rsp.Name, nil
}
