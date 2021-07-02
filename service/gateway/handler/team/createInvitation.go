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
// @Summary create invitation api
// @Description 创建 invitation
// @Tags invitation
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param expired query int false " expired 以秒为单位， 默认为 3600 秒"
// @Success 200 {object} handler.Response{data=CreateInvitationResponse}
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /team/invitation [get]
func CreateInvitation(c *gin.Context) {
	log.Info("Invitation create function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 teamID
	teamID := c.MustGet("teamID").(uint32)

	expired, err := strconv.Atoi(c.DefaultQuery("expired", "3600"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	createInvitationReq := &tpb.CreateInvitationRequest{
		TeamId:  teamID,
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
