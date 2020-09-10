package service

import (
	"fmt"
	pb "muxi-workbench-user/proto"
	"muxi-workbench/pkg/handler"

	"github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

var UserService micro.Service
var UserClient pb.UserServiceClient

// UserInit init user service
func UserInit(UserService micro.Service, UserClient pb.UserServiceClient) {
	UserService = micro.NewService(micro.Name("workbench.cli.user"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()),
	)

	UserService.Init()

	UserClient = pb.NewUserServiceClient("workbench.service.user", UserService.Client())
	fmt.Println(UserClient)
}
