package service

import (
	//tracer "muxi-workbench-status-client/tracer"
	pbu "muxi-workbench-user/proto"
	handler "muxi-workbench/pkg/handler"

	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

var UserService micro.Service
var UserClient pbu.UserServiceClient

func UserInit() {
	UserService = micro.NewService(micro.Name("workbench.cli.user"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()))
	UserService.Init()

	UserClient = pbu.NewUserServiceClient("workbench.service.user", UserService.Client())

}
