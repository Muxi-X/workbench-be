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
func UpdateTeamInfo(c *gin.Context) {
	log.Info("TeamInfo update function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req UpdateTeamInfoRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 UpdateTeamInfo 请求
	updateTeamInfoReq := &tpb.UpdateTeamInfoRequest{
		TeamId:  req.TeamID,
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
