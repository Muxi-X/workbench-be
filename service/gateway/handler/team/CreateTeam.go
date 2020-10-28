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

// CreateTeam 创建新团队
func CreateTeam(c *gin.Context) {
	log.Info("Team create function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req createTeamRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 createTeam 请求
	createTeamReq := &tpb.CreateTeamRequest{
		TeamName:  req.TeamName,
		CreatorId: req.CreatorID,
	}

	// 向 CreateTeam 服务发送请求
	_, err := service.TeamClient.CreateTeam(context.Background(), createTeamReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
