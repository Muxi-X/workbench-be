package feed

import (
	"context"
	//"fmt"
	"log"

	//"muxi-workbench-feed-client/tracer"
	pb "muxi-workbench-feed/proto"
	"muxi-workbench/pkg/handler"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

var FeedService micro.Service
var FeedClient pb.FeedServiceClient

func FeedInit(FeedService micro.Service, FeedClient pb.FeedServiceClient) {
	FeedService = micro.NewService(micro.Name("workbench.cli.feed"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()),
	)

	FeedService.Init()

	FeedClient = pb.NewFeedServiceClient("workbench.service.feed", FeedService.Client())

}

type listRequest struct {
	Role   int `json:"role"`
	Userid int `json:"userid"`
}
