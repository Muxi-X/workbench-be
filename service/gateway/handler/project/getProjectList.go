package project

import (
	"context"
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

// GetProjectList gets project list
// @Summary get project list api
// @Description 获取项目列表
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param limit query int false "limit"
// @Param last_id query int false "last_id"
// @Param page query int false "page"
// @Success 200 {object} GetProjectListResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /project/list [get]
func GetProjectList(c *gin.Context) {
	log.Info("project getProjectList function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 lastId page limit pagination

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

	// 获取 userID
	userID := c.MustGet("userID").(uint32)

	// 构造请求
	getProListReq := &pbp.GetProjectListRequest{
		UserId: userID,
		LastId: uint32(lastId),
		Offset: uint32(limit * page),
		Limit:  uint32(limit),
	}

	if page != 0 {
		getProListReq.Pagination = true
	}

	getProListResp, err := service.ProjectClient.GetProjectList(context.Background(), getProListReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	var list []*ProjectListItem
	for _, item := range getProListResp.List {
		list = append(list, &ProjectListItem{
			Id:   item.Id,
			Name: item.Name,
			Logo: item.Logo,
		})
	}

	SendResponse(c, nil, &GetProjectListResponse{
		Count: uint32(len(list)),
		List:  list,
	})
}
