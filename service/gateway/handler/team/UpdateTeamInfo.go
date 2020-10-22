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

// UpdateTeamInfo 更新团队信息
func UpdateTeamInfo(c *gin.Context) {
	log.Info("UpdateTeamInfo function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req updateTeamInfoRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 判断权限
	if req.Role != SUPERADMIN || req.Role != ADMIN {
		SendBadRequest(c, errno.ErrBind, nil, "权限不足", GetLine())
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
