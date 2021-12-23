package service

import (
	"context"
	"github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"muxi-workbench-project/model"
	"muxi-workbench/pkg/constvar"
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

func GetPath(id uint32, code uint8) string {
	path := ""

	var f func()
	f = func() { // 闭包实现递归获取path
		if code == constvar.DocFolderCode {
			folder, _ := model.GetFolderForDocDetail(id)
			path = folder.Name + "/" + path
			if folder.FatherId != 0 {
				id = folder.FatherId
				f()
			} else {
				id = folder.ProjectID
			}
		} else if code == constvar.FileFolderCode {
			folder, _ := model.GetFolderForFileDetail(id)
			path = folder.Name + "/" + path
			if folder.FatherId != 0 {
				id = folder.FatherId
				f()
			} else {
				id = folder.ProjectID
			}
		}
	}

	f()

	project, _ := model.GetProject(id)
	path = project.Name + "/" + path
	return path
}
