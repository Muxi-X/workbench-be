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

// GetApplications ... 获取申请列表
// @Summary list application api
// @Description 拉取 application 列表
// @Tags application
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param limit query int false "limit"
// @Param page query int false "page 从 0 开始计数， 如果传入非负数或者不传值则不分页"
// @Success 200 {object} ApplicationListResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /team/application/list [get]
func GetApplications(c *gin.Context) {
	log.Info("Applications list function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	pagination := true

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "-1"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}
	if page < 0 {
		pagination = false
	}

	userID := c.MustGet("userID").(uint32)

	// 构造 ApplicationList 请求
	ApplicationListReq := &tpb.ApplicationListRequest{
		Offset:     uint32(limit * page),
		Limit:      uint32(limit),
		Pagination: pagination,
		UserId:     userID,
	}

	// 向 GetApplications 服务发送请求
	listResp, err := service.TeamClient.GetApplications(context.Background(), ApplicationListReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, listResp)
}
