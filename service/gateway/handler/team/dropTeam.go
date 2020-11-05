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

// DropTeam ... 删除团队
func DropTeam(c *gin.Context) {
	log.Info("Team drop function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req DropTeamRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 DropTeam 请求
	dropTeamReq := &tpb.DropTeamRequest{
		TeamId: req.TeamID,
	}

	// 向 DropTeam 服务发送请求
	_, err := service.TeamClient.DropTeam(context.Background(), dropTeamReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}