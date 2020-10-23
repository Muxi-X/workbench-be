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

// 只用调用一次 getfiletree
// 不需要从 token 获取 userid
func GetFileTree(c *gin.Context) {
	log.Info("project getFileTree function call.",
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
	getFileTreeResp, err := service.ProjectClient.GetFileTree(context.Background(), &pbp.GetRequest{
		Id: uint32(pid),
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回
	resp := getFileTreeResponse{
		Filetree: getFileTreeResp.Tree,
	}

	SendResponse(c, nil, resp)
}
