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

// 不需要从 token 获取 userid
func GetDocTree(c *gin.Context) {
	log.Info("project getDoctTree function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 pid
	var pid int
	var err error

	pid, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 发送请求
	getDocTreeResp, err := service.ProjectClient.GetDocTree(context.Background(), &pbp.GetRequest{
		Id: uint32(pid),
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 返回结果
	resp := getDocTreeResponse{
		Doctree: getDocTreeResp.Tree,
	}

	SendResponse(c, nil, resp)
}
