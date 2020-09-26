package status

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	pbf "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/pkg/token"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbs "muxi-workbench-status/proto"

	"github.com/gin-gonic/gin"
)

// 需要调用 feed push 和 status update
// 需要从 token 获取 userid
func Update(c *gin.Context) {
	log.Info("Status update function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 sid 和请求
	var sid int
	var err error

	sid, err = strconv.Atoi(c.Param("sid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取请求
	var req updateRequest
	if err := c.Bind(&req); err != nil {
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

	// 构造 updateReq 并发送请求
	updateReq := &pbs.UpdateRequest{
		Id:      uint32(sid),
		Title:   req.Title,
		Content: req.Content,
		UserId:  uint32(ctx.ID),
	}

	_, err2 := service.StatusClient.Update(context.Background(), updateReq)
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err2.Error(), GetLine())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "编辑",
		UserId: uint32(ctx.ID),
		Source: &pbf.Source{
			Kind:        6,
			Id:          uint32(sid), // 暂时从前端获取
			Name:        req.Title,
			ProjectId:   0,
			ProjectName: "",
		},
	}

	// 向 feed 发送请求
	_, err3 := service.FeedClient.Push(context.Background(), pushReq)
	if err3 != nil {
		SendError(c, errno.InternalServerError, nil, err3.Error(), GetLine())
		return
	}

	// 返回结果
	SendResponse(c, errno.OK, nil)
}
