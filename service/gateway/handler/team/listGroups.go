package team

import (
	"context"
	"strconv"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	tpb "muxi-workbench-team/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetGroupList ... 列举组别
func GetGroupList(c *gin.Context) {
	log.Info("Groups list function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var limit int
	var page int
	var err error
	pagination := true

	limit, err = strconv.Atoi(c.Query("limit"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}
	if page == -1 {
		pagination = false
	}

	// 构造 GroupList 请求
	GroupListReq := &tpb.GroupListRequest{
		Offset:     uint32(page * limit),
		Limit:      uint32(limit),
		Pagination: pagination,
	}

	// 向 GetGroupList 服务发送请求
	listResp, err := service.TeamClient.GetGroupList(context.Background(), GroupListReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	var resp GroupListResponse
	for _, item := range listResp.List {
		resp.Groups = append(resp.Groups, Group{
			ID:        item.Id,
			Name:      item.Name,
			UserCount: item.UserCount,
		})
	}
	resp.Count = listResp.Count

	SendResponse(c, nil, resp)
}
