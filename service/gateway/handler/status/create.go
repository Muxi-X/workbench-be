package status

import (
	"context"

	pbf "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbs "muxi-workbench-status/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Create create new status
func Create(c *gin.Context) {
	log.Info("Status create function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获得请求
	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userId
	userId := c.MustGet("userID").(uint32)

	// 构造 create 请求
	createReq := &pbs.CreateRequest{
		Title:   req.Title,
		Content: req.Content,
		UserId:  userId,
	}

	// 向创建进度发起请求
	createResp, err := service.StatusClient.Create(context.Background(), createReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "创建",
		UserId: userId,
		Source: &pbf.Source{
			Kind:        6,
			Id:          createResp.Id,
			Name:        req.Title,
			ProjectId:   0,
			ProjectName: "",
		},
	}

	// 向 feed 发送请求
	_, err = service.FeedClient.Push(context.Background(), pushReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
