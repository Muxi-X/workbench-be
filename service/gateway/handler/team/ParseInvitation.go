package team

import (
	"context"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	tpb "muxi-workbench-team/proto"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// ParseInvitation 解析团队邀请码
func ParseInvitation(c *gin.Context) {
	log.Info("CreateInvitation function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 在路径中获取参数 hash
	hash := c.Param("hash")
	if hash == "" {
		SendError(c, errno.ErrBind, nil, "路径中没有读取到hash", GetLine())
		return
	}

	parseInvitationReq := &tpb.ParseInvitationRequest{Hash: hash}

	// 向 ParseInvitation 服务发送请求
	parseInvitationResp, err := service.TeamClient.ParseInvitation(context.Background(), parseInvitationReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	var resp parseInvitationResponse
	resp.TeamID = parseInvitationResp.TeamId

	SendResponse(c, nil, resp)
}
