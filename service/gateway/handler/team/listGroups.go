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

// GetGroupList 列举组别
func GetGroupList(c *gin.Context) {
	log.Info("Group list function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req updateGroupInfo
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	var limit int
	var page int
	var err error

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 GroupList 请求
	GroupListReq := &tpb.GroupListRequest{
		Offset:     uint32(page * limit),
		Limit:      uint32(limit),
		Pagination: false,
	}

	// 向 GetGroupList 服务发送请求
	listResp, err := service.TeamClient.GetGroupList(context.Background(), GroupListReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	var resp groupListResponse
	for i := 0; i < len(listResp.List); i++ {
		resp.Groups = append(resp.Groups, group{
			ID:        listResp.List[i].Id,
			Name:      listResp.List[i].Name,
			UserCount: listResp.List[i].UserCount,
		})
	}
	resp.Count = listResp.Count

	SendResponse(c, nil, resp)
}
