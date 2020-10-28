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

// GetApplications 获取申请列表
func GetApplications(c *gin.Context) {
	log.Info("Applications list function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req updateGroupInfoRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	var limit int
	var page int
	var err error
	pagination := true

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	page, err = strconv.Atoi(c.Query("page"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
	if page == -1 {
		pagination = false
	}

	// 构造 ApplicationList 请求
	ApplicationListReq := &tpb.ApplicationListRequest{
		Offset:     uint32(limit * page),
		Limit:      uint32(limit),
		Pagination: pagination,
	}

	// 向 GetApplications 服务发送请求
	listResp, err := service.TeamClient.GetApplications(context.Background(), ApplicationListReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	var resp applicationListResponse
	for i := 0; i < len(listResp.List); i++ {
		resp.ApplyList = append(resp.ApplyList, applyUserItem{
			ID:    listResp.List[i].Id,
			Name:  listResp.List[i].Name,
			Email: listResp.List[i].Email,
		})
	}
	resp.Count = listResp.Count

	SendResponse(c, nil, resp)
}
