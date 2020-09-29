package status

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/pkg/token"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbs "muxi-workbench-status/proto"

	"github.com/gin-gonic/gin"
)

// 需要调用 status like
// userid 从 token 获取
func Like(c *gin.Context) {
	log.Info("Status like function call",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取sid
	var sid int
	var err error

	sid, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userid
	raw, ifexists := c.Get("context")
	if !ifexists {
		SendBadRequest(c, errno.ErrTokenInvalid, nil, "Context not exists", GetLine())
		return
	}
	ctx, ok := raw.(*token.Context)
	if !ok {
		SendError(c, errno.ErrValidation, nil, "Context assign failed", GetLine())
		return
	}

	// 调用 like 请求
	_, err2 := service.StatusClient.Like(context.Background(), &pbs.LikeRequest{
		Id:     uint32(sid),
		UserId: uint32(ctx.ID),
	})
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err2.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
