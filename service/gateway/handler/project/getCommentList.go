package project

import (
	"context"
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

// ListComment ... 获取文档/文件评论列表
// @Summary list doc(1) or file(2) comments api
// @Description 一次获取文档/文件一二级评论列表，kind 为 1代表二级评论，一级评论在前，count为一级评论数
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param id path int true "target_id"
// @Param limit query int false "limit"
// @Param last_id query int false "last_id"
// @Param page query int false "page"
// @Param object body ListCommentRequest true "list_comment_request"
// @Param project_id query int true "project_id"
// @Success 200 {object} CommentListResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /file/comments/{id} [get]
func ListComment(c *gin.Context) {
	log.Info("Project listComment function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从 Query Param 中获取 lastId 和 limit
	targetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	lastId, err := strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取 page
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取请求体
	var req ListCommentRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 listComment 请求并发送
	listComReq := &pbp.GetCommentRequest{
		TargetId: uint32(targetID),
		Offset:   uint32(page * limit),
		Limit:    uint32(limit),
		LastId:   uint32(lastId),
		TypeId:   req.TypeId,
	}

	listComResp, err := service.ProjectClient.GetCommentList(context.Background(), listComReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	resp := CommentListResponse{
		Count: listComResp.Count,
		Total: uint32(len(listComResp.List)),
	}

	for _, item := range listComResp.List {
		resp.CommentList = append(resp.CommentList, Comment{
			Id:       item.Id,
			Uid:      item.UserId,
			UserName: item.UserName,
			Avatar:   item.Avatar,
			Time:     item.Time,
			Content:  item.Content,
			Kind:     item.Kind,
		})
	}

	SendResponse(c, nil, resp)
}
