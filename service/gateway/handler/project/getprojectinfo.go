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

// 只调用一次 project info
func GetProjectInfo(c *gin.Context) {
	log.Info("Project Info get function call")

	var pid int
	var err error

	// 获取 Pid
	pid, err = strconv.Atoi(c.Param("pid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 发送请求
	getProInfoResp, err2 := ProjectClient.GetProjectInfo(context.Background(), &pbp.GetRequest{
		Id: pid,
	})
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 构造返回 response
	resp := getProInfoResponse{
		Projectid:   getProInfoResp.Id,
		Projectname: getProInfoResp.Name,
		Intro:       getProInfoResp.Intro,
		Usercount:   getProInfoResp.Usercount,
	}

	// 返回结果
	SendResponse(c, nil, resp)
}
