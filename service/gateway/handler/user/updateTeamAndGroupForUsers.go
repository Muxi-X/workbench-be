package user

import (
	"context"
	// "strconv"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	pb "muxi-workbench-user/proto"

	"go.uber.org/zap"

	// "muxi-workbench-gateway/pkg/token"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"

	"github.com/gin-gonic/gin"
)

// UpdateTeamAndGroupForUsers ... 通过 teamid 或 groupid 给 users 数组分组/团队
func UpdateTeamAndGroupForUsers(c *gin.Context) {
	log.Info("User getInfo function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取请求
	var req UpdateTeamGroupRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造请求给 updateTeamGroup
	updateTeamGroupReq := &pb.UpdateTeamGroupRequest{
		Value: req.Value,
		Kind:  req.Kind,
	}
	for i := 0; i < len(req.Ids); i++ {
		updateTeamGroupReq.Ids = append(updateTeamGroupReq.Ids, req.Ids[i])
	}

	// 发送请求
	_, err := service.UserClient.UpdateTeamAndGroupForUsers(context.Background(), updateTeamGroupReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
