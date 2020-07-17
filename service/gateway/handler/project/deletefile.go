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

// 调用 deletefile 和 feed push
func DeleteFile(c *gin.Context) {
	log.Info("File delete function call.")

	// 获取 fid
	var fid int
	var err error

	fid, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 获取 req
	var req deleteFileRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 请求
	_, err2 := ProjectClient.DeleteFile(context.Background(), &pbp.GetRequest{
		Id: fid,
	})
	if err2 != nil {
		log.Fatal("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 构造 push 请求
	// 待确认，file 的传法
	pushReq := &pbf.PushRequest{
		Action: "删除",
		UserId: req.UserId,
		Source: &pbf.Source{
			Kind:        4,
			Id:          fid, // 暂时从前端获取
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
