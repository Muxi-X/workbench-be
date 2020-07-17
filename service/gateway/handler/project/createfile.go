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

// 调用 createfile 和 feedpush
func CreateFile(c *gin.Context) {
	log.Info("File create function call.")

	// 获取请求
	var req createFileRequest
	if err := c.BindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 构造请求
	createFileReq := &pbp.CreateFileRequest{
		ProjectId: req.Projectid,
		Name:      req.Filetname,
		HashName:  req.Hashname,
		Url:       req.Url,
		UserId:    req.UserId,
	}
	_, err2 := ProjectClient.CreateFile(context.Background(), &req)
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
			Kind:        4,
			Id:          req.Fid, // 暂时从前端获取
			Name:        req.Filename,
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
