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
	var getDetail func(uint32) (*model.FolderDetail, error)

	switch code {
	case constvar.DocFolderCode:
		getDetail = model.GetFolderForDocDetail
	case constvar.FileFolderCode:
		getDetail = model.GetFolderForFileDetail
	default:
		return "/"
	}

	path := ""
	var f func()
	f = func() { // 闭包实现递归获取path
		folder, _ := getDetail(id)
		path = folder.Name + "/" + path
		if folder.FatherId != 0 {
			id = folder.FatherId
			f()
		} else {
			id = folder.ProjectID
		}
	}

	f()

	project, _ := model.GetProject(id)
	path = project.Name + "/" + path
	return path
}
