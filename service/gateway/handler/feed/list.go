package feed

import (
	"context"
	"strconv"

	pb "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// List ... lists feeds.
// @Summary get feed list api
// @Description 获取此用户的动态list
// @Tags feed
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param limit query int false "limit"
// @Param last_id query int false "last_id"
// @Success 200 {object} FeedListResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /feed/list [get]
func List(c *gin.Context) {
	log.Info("feed List function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 limit
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取 lastId
	lastId, err := strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取 userID 和 role
	userId := c.MustGet("userID").(uint32)
	role := c.MustGet("role").(uint32)

	listReq := &pb.ListRequest{
		LastId: uint32(lastId),
		Limit:  uint32(limit),
		Role:   role,
		UserId: userId,
		Filter: &pb.Filter{
			UserId:  0,
			GroupId: 0,
		},
	}

	listResp, err := service.FeedClient.List(context.Background(), listReq)
	if err != nil {
		SendError(c, errno.ErrFeedList, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, listResp)
}
