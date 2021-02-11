package project

import (
	"context"
	pbf "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateDocComment ... 创建文档评论
func CreateDocComment(c *gin.Context) {
	log.Info("project createDocComment function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req CreateDocCommentRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	docId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	userID := c.MustGet("userID").(uint32)

	createDocCommentReq := &pbp.CreateDocCommentRequest{
		UserId:  userID,
		DocId:   uint32(docId),
		Content: req.Content,
	}

	_, err = service.ProjectClient.CreateDocComment(context.Background(), createDocCommentReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// feed
	pushReq := &pbf.PushRequest{
		Action: "评论",
		UserId: userID,
		Source: &pbf.Source{
			Kind:        3,
			Id:          uint32(docId),
			Name:        "",
			ProjectId:   req.ProjectId,
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
