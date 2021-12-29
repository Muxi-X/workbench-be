package status

import (
	"context"
	"strconv"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbs "muxi-workbench-status/proto"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CommentList ... 获取评论
// @Summary get comments list api
// @Description 通过 status_id 一次获取其一二级评论列表，kind 为 1代表二级评论，一级评论在前，count为一级评论数
// @Tags status
// @Accept  application/json
// @Produce  application/json
// @Param id path int true "status_id"
// @Param Authorization header string true "token 用户令牌"
// @Param limit query int false "limit"
// @Param last_id query int false "last_id"
// @Param page query int false "page"
// @Success 200 {object} CommentListResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /status/comment/{id} [get]
func CommentList(c *gin.Context) {
	log.Info("Status commentList function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	sid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 从 Query Param 中获取 lastId 和 limit
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

	// 构造 listComment 请求并发送
	listComReq := &pbs.CommentListRequest{
		StatusId: uint32(sid),
		Offset:   uint32(page * limit),
		Limit:    uint32(limit),
		LastId:   uint32(lastId),
	}

	listComResp, err := service.StatusClient.ListComment(context.Background(), listComReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	resp := CommentListResponse{
		Total: uint32(len(listComResp.List)),
		Count: listComResp.Count,
	}

	for _, item := range listComResp.List {
		resp.CommentList = append(resp.CommentList, Comment{
			Id:       item.Id,
			Uid:      item.UserId,
			Username: item.UserName,
			Avatar:   item.Avatar,
			Time:     item.Time,
			Content:  item.Content,
			Kind:     item.Kind,
		})
	}

	SendResponse(c, nil, resp)
}
