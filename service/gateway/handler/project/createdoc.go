package handler

import (
	"context"
	//"fmt"
	"log"

	//tracer "muxi-workbench-project-client/tracer"
	pbf "muxi-workbench-feed/proto"
	pbp "muxi-workbench-project/proto"
	handler "muxi-workbench/pkg/handler"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

// 调用一次 doc create 和 feed push
func CreateDoc(c *gin.Context) {
	log.Info("Doc create function call.")

	// 获得请求
	var req createDorRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	createDocReq := &pbp.CreateDocRequest{
		Title:     req.Title,
		Content:   req.Content,
		ProjectId: req.Pid,
		UserId:    req.UserId,
	}
	_, err2 := ProjectClient.CreateDoc(context.Background(), createDocReq)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "创建",
		UserId: req.UserId,
		Source: &pbf.Source{
			Kind:        3,
			Id:          0, // 暂时从前端获取
			Name:        req.Docname,
			ProjectId:   0,
			ProjectName: req.Pid,
		},
	}

	// 向 feed 发送请求
	_, err3 := feed.FeedClient.Push(context.Background(), pushReq)
	if err3 != nil {
		log.Fatalf("Could not greet: %v", err3)
		SendError(c, errno.InternalServerError, nil, err3.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}
