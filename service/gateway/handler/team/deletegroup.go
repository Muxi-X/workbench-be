package team

import (
	"context"
	"muxi-workbench-gateway/service"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/util"
	tpb "muxi-workbench-team/proto"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

// DeleteGroup 删除团队
func DeleteGroup(c *gin.Context) {
	log.Info("Group delete function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req deleteGroupRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 判断权限
	if req.Role != SUPERADMIN || req.Role != ADMIN {
		SendBadRequest(c, errno.ErrBind, nil, "权限不足", GetLine())
		return
	}

	// 构造 deleteGroup 请求
	deleteGroupReq := &tpb.DeleteGroupRequest{GroupId: req.GroupID}

	// 向 DeleteGroup 服务发送请求
	_, err := service.TeamClient.DeleteGroup(context.Background(), deleteGroupReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
