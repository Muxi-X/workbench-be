package team

import (
	"context"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	tpb "muxi-workbench-team/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Remove ... 移除成员
func Remove(c *gin.Context) {
	log.Info("Remove team function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req RemoveRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	removeReq := &tpb.RemoveRequest{
		UserList: req.UserList,
		TeamId:   req.TeamID,
	}

	// 向 Remove 服务发送请求
	_, err := service.TeamClient.Remove(context.Background(), removeReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
