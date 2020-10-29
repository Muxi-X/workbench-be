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

// Join ... 加入团队
func Join(c *gin.Context) {
	log.Info("Join team function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req JoinRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 join 请求
	joinReq := &tpb.JoinRequest{
		UserList: req.UserList,
		TeamId:   req.TeamID,
	}

	// 向 Join 服务发送请求
	_, err := service.TeamClient.Join(context.Background(), joinReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
