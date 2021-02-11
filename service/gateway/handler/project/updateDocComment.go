package project

import (
	"context"
	pbf "muxi-workbench-feed/proto"
	"strconv"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UpdateDocComment ... 修改文档评论
func UpdateDocComment(c *gin.Context) {
	log.Info("Project updateDocComment function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 commentID 和请求
	var err error
	var commentID int

	commentID, err = strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
	}

	docId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	userID := c.MustGet("userID").(uint32)

	// 获取请求体
	var req UpdateDocCommentRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 updateReq 并发送请求
	updateReq := &pbp.UpdateDocCommentRequest{
		Id:      uint32(commentID),
		Content: req.Content,
	}

	_, err = service.ProjectClient.UpdateDocComment(context.Background(), updateReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "编辑",
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
