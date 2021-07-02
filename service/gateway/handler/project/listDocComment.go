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

// ListDocComment ... 获取文档评论列表
// @Summary list doc comments api
// @Description 获取文档评论列表
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param id path int true "doc_id"
// @Param limit query int false "limit"
// @Param last_id query int false "last_id"
// @Param page query int false "page"
// @Param project_id query int true "project_id"
// @Success 200 {object} handler.Response{data=CommentListResponse}
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /file/doc/{id}/comments [get]
func ListDocComment(c *gin.Context) {
	log.Info("Project listDocComment function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从 Query Param 中获取 lastId 和 limit
	var limit int
	var lastId int
	var docID int
	var page int
	var err error

	docID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	lastId, err = strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取 page
	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 构造 listcomment 请求并发送
	listComReq := &pbp.ListDocCommentRequest{
		DocId:  uint32(docID),
		Offset: uint32(page * limit),
		Limit:  uint32(limit),
		LastId: uint32(lastId),
	}

	listComResp, err := service.ProjectClient.ListDocComment(context.Background(), listComReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	resp := CommentListResponse{
		Count: listComResp.Count,
	}

	for _, item := range listComResp.List {
		resp.CommentList = append(resp.CommentList, Comment{
			Cid:      item.Id,
			Uid:      item.UserId,
			Username: item.UserName,
			Avatar:   item.Avatar,
			Time:     item.Time,
			Content:  item.Content,
		})
	}

	SendResponse(c, nil, resp)
}
