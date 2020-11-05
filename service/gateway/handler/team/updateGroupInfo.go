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

// UpdateGroupInfo ... 更新组别内信息
func UpdateGroupInfo(c *gin.Context) {
	log.Info("GroupInfo update function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req UpdateGroupInfoRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 构造 updateGroupInfo 请求
	updateGroupInfoReq := &tpb.UpdateGroupInfoRequest{
		GroupId: uint32(groupID),
		NewName: req.NewGroupName,
	}

	// 向 UpdateGroupInfo 服务发送请求
	_, err = service.TeamClient.UpdateGroupInfo(context.Background(), updateGroupInfoReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
