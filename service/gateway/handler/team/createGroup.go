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

// CreateGroup 创建组别
func CreateGroup(c *gin.Context) {
	log.Info("Group create function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req createGroupRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 createGroup 请求
	createGroupReq := &tpb.CreateGroupRequest{
		GroupName: req.GroupName,
		UserList:  req.UserList,
	}

	// 向 CreateGroup 服务发送请求
	_, err := service.TeamClient.CreateGroup(context.Background(), createGroupReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
