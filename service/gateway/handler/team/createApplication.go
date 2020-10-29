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

// CreateApplication ... 创建申请
func CreateApplication(c *gin.Context) {
	log.Info("Application create function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req ApplicationRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 CreateApplication 请求
	CreateApplicationReq := &tpb.ApplicationRequest{
		UserId: req.UserID,
	}

	// 向 CreateApplication 服务发送请求
	_, err := service.TeamClient.CreateApplication(context.Background(), CreateApplicationReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
