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

// 只用调用一次 getfiletree
func GetFileTree(c *gin.Context) {
	log.Info("Project filetree get function call.")

	// 获取 pid
	var pid int
	var err error

	pid, err = strconv.Atoi(c.Param("pid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 发送请求
	getFileTreeResp, err2 := ProjectClient.GetFileTree(context.Background(), &pbp.GetRequest{
		Id: pid,
	})
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 构造返回
	resp := getFileTreeResponse{
		Filetree: getFileTreeResp.Filetree,
	}

	SendResponse(c, nil, resp)
}
