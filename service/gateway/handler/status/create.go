package status

import (
	"context"
	//"fmt"
	"log"

	//tracer "muxi-workbench-status-client/tracer"
	pbf "muxi-workbench-feed/proto"
	"muxi-workbench-gateway/handler/feed"
	pbs "muxi-workbench-status/proto"
	handler "muxi-workbench/pkg/handler"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
)

func Create(c *gin.Context) {
	log.Info("Status create function call.")

	// 获得请求
	var req createRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 构造 create 请求
	createReq := &pbs.CreateRequest{
		Title:   req.Title,
		Content: req.Content,
		UserId:  req.UserId,
	}

	// 向创建进度服务发送请求
	_, err := StatusClient.Create(context.Background(), createReq)

	if err != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "创建",
		UserId: req.UserId,
		Source: &pbf.Source{
			Kind:        6,
			Id:          req.Statusid, // 暂时从前端获取
			Name:        req.Title,
			ProjectId:   0,
			ProjectName: "",
		},
	}

	// 向 feed 发送请求
	_, err2 := feed.FeedClient.Push(context.Background(), pushReq)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err2.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}
