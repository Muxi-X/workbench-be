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

// DeleteDocComment ... 删除文档评论
func DeleteDocComment(c *gin.Context) {
	log.Info("project deleteDocComment function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	docId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	var req DeleteDocCommentRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	userID := c.MustGet("userID").(uint32)

	deleteDocCommentReq := &pbp.DeleteDocCommentRequest{
		UserId:    userID,
		CommentId: uint32(commentId),
	}

	_, err = service.ProjectClient.DeleteDocComment(context.Background(), deleteDocCommentReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// feed
	pushReq := &pbf.PushRequest{
		Action: "取消评论",
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
