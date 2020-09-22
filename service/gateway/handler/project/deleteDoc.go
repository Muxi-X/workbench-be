package project

import (
	"context"
	"strconv"

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
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
)

// 调用一次 deletedoc 和一次 feed push
<<<<<<< HEAD
// 需要从 token 获取 userid
=======
>>>>>>> master
func DeleteDoc(c *gin.Context) {
	log.Info("Doc delete function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 did
	var did int
	var err error

	did, err = strconv.Atoi(c.Param("did"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 获取请求
	var req deleteDocRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

<<<<<<< HEAD
	// 获取 userid
	raw, ifexists := c.Get("context")
	if !ifexists {
		SendBadRequest(c, errno.ErrTokenInvalid, nil, "Context not exists")
	}
	ctx, ok := raw.(*token.Context)
	if !ok {
		SendError(c, errno.ErrValidation, nil, "Context assign failed")
	}

=======
>>>>>>> master
	_, err2 := service.ProjectClient.DeleteDoc(context.Background(), &pbp.GetRequest{
		Id: uint32(did),
	})
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "删除",
<<<<<<< HEAD
		UserId: uint32(ctx.ID),
=======
		UserId: req.UserId,
>>>>>>> master
		Source: &pbf.Source{
			Kind:        3,
			Id:          uint32(did), // 暂时从前端获取
			Name:        req.Docname,
			ProjectId:   0,
			ProjectName: "",
		},
	}

	// 向 feed 发送请求
	_, err3 := service.FeedClient.Push(context.Background(), pushReq)
	if err3 != nil {
		SendError(c, errno.InternalServerError, nil, err3.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}
