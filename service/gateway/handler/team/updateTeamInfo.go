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

// UpdateTeamInfo ... 更新团队信息
// @Summary update team info api
// @Description 更新团队信息
// @Tags team
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body UpdateTeamInfoRequest true "update_team_info_request"
// @Security ApiKeyAuth
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /team [put]
func UpdateTeamInfo(c *gin.Context) {
	log.Info("TeamInfo update function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req UpdateTeamInfoRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 teamID
	teamID := c.MustGet("teamID").(uint32)

	// 构造 UpdateTeamInfo 请求
	updateTeamInfoReq := &tpb.UpdateTeamInfoRequest{
		TeamId:  teamID,
		NewName: req.NewName,
	}

	// 向 UpdateTeamInfo 服务发送请求
	_, err := service.TeamClient.UpdateTeamInfo(context.Background(), updateTeamInfoReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
