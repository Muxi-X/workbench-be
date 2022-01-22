package status

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbs "muxi-workbench-status/proto"
	"strconv"
)

// Search ... 全文 + 标题搜索 可以对组别和用户做筛选。
// @Summary searches from title and content of status api
// @Description group_id为0时，搜索所有进度
// @Tags status
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param limit query int false "limit"
// @Param last_id query int false "last_id"
// @Param page query int false "page"
// @Param object body SearchStatusRequest true "search_status_request"
// @Success 200 {object} SearchResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /status/search [post]
func Search(c *gin.Context) {
	log.Info("project search function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 lastId page limit
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	lastId, err := strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	var req SearchStatusRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	searchReq := &pbs.SearchRequest{
		Keyword:  req.Keyword,
		UserName: req.UserName,
		Offset:   uint32(limit * page),
		LastId:   uint32(lastId),
		Limit:    uint32(limit),
		GroupId:  req.GroupID,
	}
	if page != 0 {
		searchReq.Pagination = true
	}

	resp, err := service.StatusClient.Search(context.Background(), searchReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	var list []*SearchResult
	for _, v := range resp.List {
		list = append(list, &SearchResult{
			Id:       v.Id,
			Title:    v.Title,
			Content:  v.Content,
			UserName: v.UserName,
			Time:     v.Time,
		})
	}

	SendResponse(c, errno.OK, &SearchResponse{
		List:  list,
		Count: resp.Count,
	})
}
