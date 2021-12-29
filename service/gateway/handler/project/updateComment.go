package project

import (
	"context"
	pbf "muxi-workbench-feed/proto"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"
)

// UpdateComment ... 修改文档/文件评论
// @Summary update doc or file comment api
// @Description 修改文档/文件评论
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param id path int true "target_id"
// @Param comment_id path int true "comment_id"
// @Param object body UpdateCommentRequest true "update_comment_request"
// @Param project_id query int true "此文档/文件所属项目 id"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /file/comment/{id}/{comment_id} [put]
func UpdateComment(c *gin.Context) {
	log.Info("Project updateComment function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 commentID 和请求
	commentID, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
	}

	targetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	userID := c.MustGet("userID").(uint32)

	// 获取请求体
	var req UpdateCommentRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 updateReq 并发送请求
	updateReq := &pbp.UpdateCommentRequest{
		CommentId: uint32(commentID),
		Content:   req.Content,
		UserId:    userID,
		TypeId:    req.TypeId,
	}

	_, err = service.ProjectClient.UpdateComment(context.Background(), updateReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	projectID := c.MustGet("projectID").(uint32)

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "编辑",
		UserId: userID,
		Source: &pbf.Source{
			Kind:      3, // TODO
			Id:        uint32(targetId),
			Name:      "",
			ProjectId: projectID,
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
