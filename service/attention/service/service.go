package service

import (
	"context"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"muxi-workbench-attention/model"
	"muxi-workbench/pkg/handler"

	ppb "muxi-workbench-project/proto"
	upb "muxi-workbench-user/proto"

	"github.com/micro/go-micro"
)

var UserService micro.Service
var UserClient upb.UserServiceClient
var ProjectService micro.Service
var ProjectClient ppb.ProjectServiceClient

func UserInit() {
	UserService = micro.NewService(micro.Name("workbench.cli.user"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()),
	)

	UserService.Init()

	UserClient = upb.NewUserServiceClient("workbench.service.user", UserService.Client())
}

func ProjectInit() {
	ProjectService = micro.NewService(micro.Name("workbench.cli.project"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()),
	)

	ProjectService.Init()

	ProjectClient = ppb.NewProjectServiceClient("workbench.service.project", ProjectService.Client())
}

// AttentionService ... 动态服务
type AttentionService struct{}

// GetDocFromProjectService get doc data from project-service
func GetDocFromProjectService(id uint32) (*model.File, error) {
	rsp, err := ProjectClient.GetDocDetail(context.Background(), &ppb.GetFileDetailRequest{Id: id})
	if err != nil {
		return nil, err
	}

	doc := &model.File{
		CreatorName: rsp.Creator,
		Name:        rsp.Title,
		Id:          rsp.Id,
		Kind:        0,
		// ProjectId:   rsp. TODO:project新增一个通过doc获取project rpc
	}
	if err != nil {
		return doc, err
	}
	return doc, nil
}

// GetFileFromProjectService get file data from project-service
func GetFileFromProjectService(id uint32) (*model.File, error) {
	// rsp, err := ProjectClient.GetProjectIdsForUser(context.Background(), &ppb.GetRequest{Id: id})
	rsp, err := ProjectClient.GetFileDetail(context.Background(), &ppb.GetFileDetailRequest{Id: id})
	if err != nil {
		return nil, err
	}

	doc := &model.File{
		CreatorName: rsp.Creator,
		Name:        rsp.Name,
		Id:          rsp.Id,
		Kind:        1,
		// ProjectId:   rsp. TODO:project新增一个通过doc获取project rpc
	}
	if err != nil {
		return doc, err
	}
	return doc, nil
}

// GetInfoFromUserService get user's name and avatar from user-service
func GetInfoFromUserService(id uint32) (string, error) {
	rsp, err := UserClient.GetProfile(context.Background(), &upb.GetRequest{Id: id})
	if err != nil {
		return "", err
	}

	return rsp.Name, nil
}
