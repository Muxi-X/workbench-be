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

// 需要调用 update 和 feed push
func UpdateProjectInfo(c *gin.Context) {
	log.Info("Project info update function call")

	// 获取 pid
	var pid int
	var err error

	pid, err = strconv.Atoi(c.Param("pid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 获取请求
	var req updateRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 构造 update 请求
	updateReq := &pbp.ProjectInfo{
		Id:        pid,
		Name:      req.Projectname,
		intro:     req.Intro,
		UserCount: req.Usercount,
	}

	// 发送 update 请求
	_, err2 := ProjectClient.UpdateProjectInfo(context.Background(), updateReq)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err2.Error())
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
			ProjectName: req.Projectname,
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
