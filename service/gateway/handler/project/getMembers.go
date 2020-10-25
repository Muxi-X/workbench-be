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

// GetMembers gets the members in the project
func GetMembers(c *gin.Context) {
	log.Info("Project getMembers function call.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 要从 query param 获取 lastid limit page pagination

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	lastId, err := strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	pagination, err := strconv.Atoi(c.DefaultQuery("pagination", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 projectID
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	var pageBool bool
	if pagination == 1 {
		pageBool = true
	}

	// 构造请求
	getMembersRequest := &pbp.GetMemberListRequest{
		ProjectId:  uint32(projectID),
		Lastid:     uint32(lastId),
		Offset:     uint32(page * limit),
		Limit:      uint32(limit),
		Pagination: pageBool,
	}

	getMembersResponse, err := service.ProjectClient.GetMembers(context.Background(), getMembersRequest)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	var list []*MemberListItem
	for _, item := range getMembersResponse.List {
		list = append(list, &MemberListItem{
			Id:        item.Id,
			Name:      item.Name,
			Avatar:    item.Avatar,
			GroupName: item.GroupName,
			Role:      item.Role,
		})
	}

	SendResponse(c, nil, &GetMemberResponse{
		Count: getMembersResponse.Count,
		List:  list,
	})
}
