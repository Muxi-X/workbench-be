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

// DeleteGroup 删除团队
func DeleteGroup(c *gin.Context) {
	log.Info("Group delete function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 deleteGroup 请求
	deleteGroupReq := &tpb.DeleteGroupRequest{GroupId: uint32(groupID)}

	// 向 DeleteGroup 服务发送请求
	_, err = service.TeamClient.DeleteGroup(context.Background(), deleteGroupReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
