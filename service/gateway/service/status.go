package service

import (
	//tracer "muxi-workbench-status-client/tracer"
	pbs "muxi-workbench-status/proto"
	handler "muxi-workbench/pkg/handler"

	micro "github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

var StatusService micro.Service
var StatusClient pbs.StatusServiceClient

func StatusInit() {
	StatusService = micro.NewService(micro.Name("workbench.cli.status"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()))
	StatusService.Init()

	StatusClient = pbs.NewStatusServiceClient("workbench.service.status", StatusService.Client())

}
