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

// ListGroup ... list feeds filtered by group id.
// 0 代表不筛选，1->产品，2->前端，3->后端，4->安卓，5->设计
// @Summary get feed group_list by group_id api
// @Description 获取某一组的动态list，0 代表不筛选，1->产品，2->前端，3->后端，4->安卓，5->设计
// @Tags feed
// @Accept  application/json
// @Produce  application/json
// @Param id path int true "group_id"
// @Param Authorization header string true "token 用户令牌"
// @Param limit query int false "limit"
// @Param last_id query int false "last_id"
// @Success 200 {object} FeedListResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /feed/list/group/{id} [get]
func ListGroup(c *gin.Context) {
	log.Info("Feed listGroup function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 groupId
	groupId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	lastId, err := strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取 userId 和 role
	userId := c.MustGet("userID").(uint32)
	role := c.MustGet("role").(uint32)

	listReq := &pb.ListRequest{
		LastId: uint32(lastId),
		Limit:  uint32(limit),
		Role:   role,
		UserId: userId,
		Filter: &pb.Filter{
			UserId:  0,
			GroupId: uint32(groupId),
		},
	}

	listResp, err := service.FeedClient.List(context.Background(), listReq)
	if err != nil {
		SendError(c, errno.ErrFeedList, nil, err.Error(), GetLine())
		return
	}

	var list []*FeedItem
	for _, item := range listResp.List {
		list = append(list, &FeedItem{
			Id:          item.Id,
			Action:      item.Action,
			ShowDivider: item.ShowDivider,
			Date:        item.Date,
			Time:        item.Time,
			FeedUser: &FeedUser{
				Name:      item.User.Name,
				Id:        item.User.Id,
				AvatarUrl: item.User.AvatarUrl,
			},
			Source: &Source{
				Kind:        item.Source.Kind,
				Id:          item.Source.Id,
				Name:        item.Source.Name,
				ProjectId:   item.Source.ProjectId,
				ProjectName: item.Source.ProjectName,
			},
		})
	}

	SendResponse(c, nil, FeedListResponse{
		Count: listResp.Count,
		List:  list,
	})
}
