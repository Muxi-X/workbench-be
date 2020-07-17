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

// 调用一次 deletedoc 和一次 feed push
func DeleteDoc(c *gin.Context) {
	log.Info("Doc delete function call.")

	// 获取 did
	var did int
	var err error

	did, err = strconv.Atoi(c.Param("did"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 获取请求
	var req pb.DeleteDocRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	_, err2 := ProjectClient.DeleteDoc(context.Background(), &pbp.GetRequest{
		Id: did,
	})
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "删除",
		UserId: req.UserId,
		Source: &pbf.Source{
			Kind:        3,
			Id:          did, // 暂时从前端获取
			Name:        req.Docname,
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
