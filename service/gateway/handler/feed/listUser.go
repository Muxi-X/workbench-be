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

// ListUser list feeds filtered by user id.
func ListUser(c *gin.Context) {
	log.Info("Feed listUser function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取筛选的 userId
	filteredUserId, err := strconv.Atoi(c.Param("id"))
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

	// 获取请求方的 userId 和 role
	userId := c.MustGet("userID").(uint32)
	role := c.MustGet("role").(uint32)

	listReq := &pb.ListRequest{
		LastId: uint32(lastId),
		Limit:  uint32(limit),
		Role:   role,
		UserId: userId,
		Filter: &pb.Filter{
			UserId:  uint32(filteredUserId),
			GroupId: 0,
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
			User: &User{
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

	SendResponse(c, nil, ListResponse{
		Count: listResp.Count,
		List:  list,
	})
}
