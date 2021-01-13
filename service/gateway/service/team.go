package service

import (
	//tracer "muxi-workbench-status-client/tracer"
	tpb "muxi-workbench-team/proto"
	handler "muxi-workbench/pkg/handler"

	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

var TeamService micro.Service
var TeamClient tpb.TeamServiceClient

func TeamInit() {
	UserService = micro.NewService(micro.Name("workbench.cli.team"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()))
	TeamService.Init()

	TeamClient = tpb.NewTeamServiceClient("workbench.service.team", TeamService.Client())
}
