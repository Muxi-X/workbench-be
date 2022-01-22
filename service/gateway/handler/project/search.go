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
	"strconv"
)

// Search ... 文档，文件分别进行搜索，文档是全文 + 标题搜索，文件是标题。可以是全局的（用户所在的所有项目）或者针对某个项目。
// @Summary searches from title and content of doc or title of file api
// @Description search_type: 为0时搜索doc，为1时搜索file（project_id为0时，搜索该用户所有项目）
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param limit query int false "limit"
// @Param last_id query int false "last_id"
// @Param page query int false "page"
// @Param object body SearchProjectRequest true "search_project_request"
// @Success 200 {object} SearchResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /project/search [post]
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

	var req SearchProjectRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userid
	userID := c.MustGet("userID").(uint32)

	searchReq := &pbp.SearchRequest{
		Type:      req.Type,
		Keyword:   req.Keyword,
		UserId:    userID,
		LastId:    uint32(lastId),
		Offset:    uint32(limit * page),
		Limit:     uint32(limit),
		ProjectId: req.ProjectID,
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
