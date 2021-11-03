package service

import (
	//tracer "muxi-workbench-attention-client/tracer"
	pb "muxi-workbench-attention/proto"
	handler "muxi-workbench/pkg/handler"

	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

var AttentionService micro.Service
var AttentionClient pb.AttentionServiceClient

func AttentionInit() {
	AttentionService = micro.NewService(
		micro.Name("workbench.cli.attention"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()),
	)

	AttentionService.Init()

	AttentionClient = pb.NewAttentionServiceClient("workbench.service.attention", AttentionService.Client())

}
