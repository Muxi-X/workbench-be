package handler

import (
	"context"
	//"fmt"
	"log"

	//tracer "muxi-workbench-project-client/tracer"
	pbp "muxi-workbench-project/proto"
	handler "muxi-workbench/pkg/handler"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

func GetDocTree(c *gin.Context) {
	log.Info("Project doctree get function call.")

	// 获取 pid
	var pid int
	var err error

	pid, err = strconv.Atoi(c.Param("pid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 发送请求
	getDocTreeResp, err2 := ProjectClient.GetDocTree(context.Background(), &pbp.GetRequest{
		Id: pid,
	})
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 返回结果
	resp := getDocTreeResponse{
		Doctree: getDocTreeResp.Tree,
	}

	SendResponse(c, nil, resp)
}
