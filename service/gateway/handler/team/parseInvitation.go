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

// ParseInvitation ... 解析团队邀请码
// @Summary parse invitation api
// @Description 解析 invitation
// @Tags invitation
// @Accept  application/json
// @Produce  application/json
// @Param hash path string true "hash"
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} ParseInvitationResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /team/invitation/{hash} [get]
func ParseInvitation(c *gin.Context) {
	log.Info("Invitation parse function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 在路径中获取参数 hash
	hash := c.Param("hash")
	if hash == "" {
		SendError(c, errno.ErrQuery, nil, "路径中没有读取到hash", GetLine())
		return
	}

	parseInvitationReq := &tpb.ParseInvitationRequest{Hash: hash}

	// 向 ParseInvitation 服务发送请求
	parseInvitationResp, err := service.TeamClient.ParseInvitation(context.Background(), parseInvitationReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	var resp ParseInvitationResponse
	resp.TeamID = parseInvitationResp.TeamId

	SendResponse(c, nil, resp)
}
