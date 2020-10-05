package service

import (
	//tracer "muxi-workbench-feed-client/tracer"
	pb "muxi-workbench-feed/proto"
	handler "muxi-workbench/pkg/handler"

	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

var FeedService micro.Service
var FeedClient pb.FeedServiceClient

func FeedInit() {
	FeedService = micro.NewService(micro.Name("workbench.cli.feed"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()),
	)

	FeedService.Init()

	FeedClient = pb.NewFeedServiceClient("workbench.service.feed", FeedService.Client())

}
