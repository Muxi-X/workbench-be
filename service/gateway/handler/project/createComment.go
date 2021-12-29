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

// CreateComment ... 创建文档/文件评论
// @Summary create a doc or file comment api
// @Description 新建文档/文件评论
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body CreateCommentRequest true "create_comment_request"
// @Param id path int true "target_id"
// @Param project_id query int true "project_id"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /file/comment/{id} [post]
func CreateComment(c *gin.Context) {
	log.Info("project createComment function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req CreateCommentRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	targetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	userID := c.MustGet("userID").(uint32)
	projectID := c.MustGet("projectID").(uint32)

	createDocCommentReq := &pbp.CreateCommentRequest{
		UserId:   userID,
		TargetId: uint32(targetId),
		TypeId:   req.TypeId,
		Content:  req.Content,
	}

	_, err = service.ProjectClient.CreateComment(context.Background(), createDocCommentReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// feed
	pushReq := &pbf.PushRequest{
		Action: "评论",
		UserId: userID,
		Source: &pbf.Source{
			Kind:      3,
			Id:        uint32(targetId),
			Name:      "", // TODO
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
