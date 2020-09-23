package service

import (
	//tracer "muxi-workbench-project-client/tracer"
	pbp "muxi-workbench-project/proto"
	handler "muxi-workbench/pkg/handler"

	micro "github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

var ProjectService micro.Service
var ProjectClient pbp.ProjectServiceClient

<<<<<<< HEAD
func ProjectInit() {
=======
func ProjectInit(ProjectService micro.Service, ProjectClient pbp.ProjectServiceClient) {
>>>>>>> master
	ProjectService = micro.NewService(micro.Name("workbench.cli.project"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()),
	)

	ProjectService.Init()

	ProjectClient = pbp.NewProjectServiceClient("workbench.service.project", ProjectService.Client())
}
