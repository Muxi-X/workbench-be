package project

import (
	"context"

	"go.uber.org/zap"
	pbf "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/pkg/token"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
)

// 调用一次 doc create 和 feed push
// 需要从 token 获取 userid
func CreateDoc(c *gin.Context) {
	log.Info("project createDoc function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获得请求
	var req createDocRequest
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

	createDocReq := &pbp.CreateDocRequest{
		Title:     req.Title,
		Content:   req.Content,
		ProjectId: req.Pid,
		UserId:    uint32(ctx.ID),
	}
	_, err2 := service.ProjectClient.CreateDoc(context.Background(), createDocReq)
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err2.Error(), GetLine())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "创建",
		UserId: uint32(ctx.ID),
		Source: &pbf.Source{
			Kind:        3,
			Id:          0, // 暂时从前端获取
			Name:        req.Docname,
			ProjectId:   req.Pid,
			ProjectName: "",
		},
	}

	// 向 feed 发送请求
	_, err3 := service.FeedClient.Push(context.Background(), pushReq)
	if err3 != nil {
		SendError(c, errno.InternalServerError, nil, err3.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
