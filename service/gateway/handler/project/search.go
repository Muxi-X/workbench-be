package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"
)

// Search
// @Summary searches from title of doc and file or content of doc api
// @Description search_type: 为0时搜索doc_title and file_title，为1时继续搜索doc_content
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body SearchRequest true "search_request"
// @Success 200 {object} SearchResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /project/search [post]
func Search(c *gin.Context) {
	log.Info("project search function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req SearchRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userid
	userID := c.MustGet("userID").(uint32)
	searchReq := &pbp.SearchRequest{
		Type:       req.Type,
		Keyword:    req.Keyword,
		UserId:     userID,
		Offset:     req.Offset,
		Limit:      req.Limit,
		Pagination: req.Pagination,
	}

	resp, err := service.ProjectClient.Search(context.Background(), searchReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	var list []*SearchResult
	for _, v := range resp.List {
		list = append(list, &SearchResult{
			Id:          v.Id,
			Title:       v.Title,
			Content:     v.Content,
			UserName:    v.UserName,
			ProjectName: v.ProjectName,
			Time:        v.Time,
		})
	}

	SendResponse(c, errno.OK, &SearchResponse{
		List:  list,
		Count: resp.Count,
	})
}
