package service

import (
	"context"
	"github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"muxi-workbench/pkg/handler"

	apb "muxi-workbench-attention/proto"
)

// Service ... 项目服务
type Service struct {
}

var AttentionClient apb.AttentionServiceClient
var AttentionService micro.Service

func AttentionInit() {
	AttentionService = micro.NewService(micro.Name("workbench.cli.attention"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()),
	)

	AttentionService.Init()

	AttentionClient = apb.NewAttentionServiceClient("workbench.service.attention", AttentionService.Client())
}

// DeleteAttentionsFromAttentionService delete linked-attentions when file was deleted from attention-service
func DeleteAttentionsFromAttentionService(id uint32, kind, userID uint32) error {
	req := &apb.PushRequest{
		UserId:   userID,
		FileId:   id,
		FileKind: kind,
	}
	_, err := AttentionClient.Delete(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}
