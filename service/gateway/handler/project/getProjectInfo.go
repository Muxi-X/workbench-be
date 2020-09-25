package project

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
)

// 只调用一次 project info
// 不需要从 token 获取 userid
func GetProjectInfo(c *gin.Context) {
	log.Info("Project Info get function call",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var pid int
	var err error

	// 获取 Pid
	pid, err = strconv.Atoi(c.Param("pid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 发送请求
	getProInfoResp, err2 := service.ProjectClient.GetProjectInfo(context.Background(), &pbp.GetRequest{
		Id: uint32(pid),
	})
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err2.Error())
		return
	}

	// 构造返回 response
	resp := getProInfoResponse{
		Projectid:   getProInfoResp.Id,
		Projectname: getProInfoResp.Name,
		Intro:       getProInfoResp.Intro,
		Usercount:   getProInfoResp.UserCount,
	}

	// 返回结果
	SendResponse(c, nil, resp)
}
