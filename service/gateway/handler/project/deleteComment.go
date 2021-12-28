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

// DeleteComment ... 删除文档/文件评论
// @Summary delete a doc or file comment api
// @Description 删除文档/文件的评论
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body DeleteCommentRequest true "delete_comment_request"
// @Param id path int true "target_id"
// @Param comment_id path int true "comment_id"
// @Param project_id query int true "project_id"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /file/comment/{id}/{comment_id} [delete]
func DeleteComment(c *gin.Context) {
	log.Info("project deleteComment function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	targetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	var req DeleteCommentRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	userID := c.MustGet("userID").(uint32)

	deleteDocCommentReq := &pbp.DeleteCommentRequest{
		TypeId:    req.TypeId,
		UserId:    userID,
		CommentId: uint32(commentId),
	}

	_, err = service.ProjectClient.DeleteComment(context.Background(), deleteDocCommentReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// feed
	pushReq := &pbf.PushRequest{
		Action: "取消评论",
		UserId: userID,
		Source: &pbf.Source{
			Kind:        3, // TODO
			Id:          uint32(targetId),
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
