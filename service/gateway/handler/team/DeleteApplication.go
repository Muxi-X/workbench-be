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

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// DeleteApplication 删除申请
func DeleteApplication(c *gin.Context) {
	log.Info("Group create function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	UserID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 DeleteApplication 请求
	DeleteApplicationReq := &tpb.ApplicationRequest{
		UserId: uint32(UserID),
	}

	// 向 DeleteApplication 服务发送请求
	_, err = service.TeamClient.DeleteApplication(context.Background(), DeleteApplicationReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
