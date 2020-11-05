package team

import (
	"context"
	"strconv"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	tpb "muxi-workbench-team/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateInvitation ... 创建团队邀请码
func CreateInvitation(c *gin.Context) {
	log.Info("Invitation create function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req CreateInvitationRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	expired, err := strconv.Atoi(c.DefaultQuery("expired", "3600"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	createInvitationReq := &tpb.CreateInvitationRequest{
		TeamId:  req.TeamID,
		Expired: int64(expired),
	}

	// 向 CreateInvitation 服务发送请求
	CreateInvitationResp, err := service.TeamClient.CreateInvitation(context.Background(), createInvitationReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	var resp = CreateInvitationResponse{Hash: CreateInvitationResp.Hash}

	SendResponse(c, nil, resp)
}
