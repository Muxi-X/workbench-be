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

// CreateTeam ... 创建新团队
// @Summary create team api
// @Description 创建 team
// @Tags team
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body CreateTeamRequest true "create_team_request"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /team [post]
func CreateTeam(c *gin.Context) {
	log.Info("Team create function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req CreateTeamRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 creatorID
	creatorID := c.MustGet("userID").(uint32)

	// 构造 createTeam 请求
	createTeamReq := &tpb.CreateTeamRequest{
		TeamName:  req.TeamName,
		CreatorId: creatorID,
	}

	// 向 CreateTeam 服务发送请求
	_, err := service.TeamClient.CreateTeam(context.Background(), createTeamReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
