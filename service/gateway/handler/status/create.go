package status

import (
	"context"

	"go.uber.org/zap"
	pbf "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
<<<<<<< HEAD
	"muxi-workbench-gateway/pkg/token"
=======
>>>>>>> master
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbs "muxi-workbench-status/proto"

	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
// userid 从 token 获取
=======
>>>>>>> master
func Create(c *gin.Context) {
	log.Info("Status create function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获得请求
	var req createRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

<<<<<<< HEAD
	// 获取 userid
	raw, ifexists := c.Get("context")
	if !ifexists {
		SendBadRequest(c, errno.ErrTokenInvalid, nil, "context not exists")
		return
	}
	ctx, ok := raw.(*token.Context)
	if !ok {
		SendError(c, errno.ErrValidation, nil, "context assign failed")
		return
	}

=======
>>>>>>> master
	// 构造 create 请求
	createReq := &pbs.CreateRequest{
		Title:   req.Title,
		Content: req.Content,
<<<<<<< HEAD
		UserId:  uint32(ctx.ID),
=======
		UserId:  req.UserId,
>>>>>>> master
	}

	// 向创建进度服务发送请求
	_, err := service.StatusClient.Create(context.Background(), createReq)

	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "创建",
<<<<<<< HEAD
		UserId: uint32(ctx.ID),
=======
		UserId: req.UserId,
>>>>>>> master
		Source: &pbf.Source{
			Kind:        6,
			Id:          req.Statusid, // 暂时从前端获取
			Name:        req.Title,
			ProjectId:   0,
			ProjectName: "",
		},
	}

	// 向 feed 发送请求
	_, err2 := service.FeedClient.Push(context.Background(), pushReq)
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err2.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}
