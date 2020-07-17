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

func UpdateDocTree(c *gin.Context) {
	log.Info("Project doctree update funcation call.")

	// 获取 pid
	var pid int
	var err error

	pid, err = strconv.Atoi(c.Param("pid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	var req updateDocTreeRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 发送请求
	_, err2 := ProjectClient.UpdateDocTree(context.Background(), &pbp.UpdateTreeRequest{
		Id:   pid,
		Tree: req.Doctree,
	})
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "编辑",
		UserId: req.UserId,
		Source: &pbf.Source{
			Kind:        2,
			Id:          0, // 暂时从前端获取
			Name:        "",
			ProjectId:   pid,
			ProjectName: req.Username,
		},
	}

	// 向 feed 发送请求
	_, err3 := feed.FeedClient.Push(context.Background(), pushReq)
	if err3 != nil {
		log.Fatalf("Could not greet: %v", err3)
		SendError(c, errno.InternalServerError, nil, err2.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}
