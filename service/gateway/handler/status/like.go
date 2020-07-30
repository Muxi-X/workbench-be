package status

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbs "muxi-workbench-status/proto"

	"github.com/gin-gonic/gin"
)

// 需要调用 status like
func Like(c *gin.Context) {
	log.Info("Status like function call",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取sid
	var sid int
	var err error

	sid, err = strconv.Atoi(c.Param("sid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 获取请求
	var req pbs.LikeRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 调用 like 请求
	_, err2 := service.StatusClient.Like(context.Background(), &pbs.LikeRequest{
		Id:     uint32(sid),
		UserId: req.UserId,
	})
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err2.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}