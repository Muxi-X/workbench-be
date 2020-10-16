package status

import (
	"context"
	"strconv"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbs "muxi-workbench-status/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Like ... 给 Status 点赞
func Like(c *gin.Context) {
	log.Info("Status like function call", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取sid
	sid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取 userID
	userID := c.MustGet("userID").(uint32)

	// 调用 like 请求
	_, err = service.StatusClient.Like(context.Background(), &pbs.LikeRequest{
		Id:     uint32(sid),
		UserId: userID,
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
