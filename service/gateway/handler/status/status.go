package status

import (
	"context"
	//"fmt"
	"log"

	//tracer "muxi-workbench-status-client/tracer"
	pbs "muxi-workbench-status/proto"
	handler "muxi-workbench/pkg/handler"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
)

var StatusService micro.Servic
var StatusClient pbs.StatusServiceClient

func StatusInit(StatusService micro.Service, StatusClient pbs.StatusServiceClient) {
	StatusService = micro.NewService(micro.Name("workbench.cli.status"),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()))
	StatusService.Init()

	StatusClient = pbs.NewStatusServiceClient("workbench.service.status", StatusService.Client())

}

type createRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserId   int    `json:"userid"`
	Statusid int    `json:"stautsid"`
}

type comment struct {
	Cid      int    `json:"cid"`
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Time     string `json:"time"`
	Content  string `json:"content"`
}

type getResponse struct {
	Sid         int       `json:"sid"`
	Title       int       `json:"title"`
	Content     string    `json:"content"`
	UserId      int       `json:userid"`
	Time        string    `json:"time"`
	Avatar      string    `json:"avatar"`
	Username    string    `json:"username"`
	Count       int       `json:"count"`
	Commentlist []comment `json:"commentlist"`
}

type updateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"userid"`
}

type deleteRequest struct {
	UserId int `json:"userid"`
	Title  int `json:title"`
}

type status struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserId   int    `json:"userid"`
	Time     string `json:"time"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
}

type listResponse struct {
	Count  int      `json:"count"`
	Status []status `json:"stauts"`
}

type likeRequest struct {
	UserId int `json:"userid"`
}

type CreateCommentRequest struct {
	Content string `json:"content"`
	UserId  int    `json:"userid"`
	Title   string `json:"title"`
}
