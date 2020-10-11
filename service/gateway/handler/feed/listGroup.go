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

// Feed 的 ListGroup 接口
func ListGroup(c *gin.Context) {
	log.Info("Feed listGroup function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 groupId
	groupId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "50"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	lastId, err := strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 role
	// TO DO: 从 Authorization 获取或从 user-service 获取
	var req ListRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userid
	id := c.MustGet("userID").(uint32)

	listReq := &pb.ListRequest{
		LastId: uint32(lastId),
		Limit:  uint32(limit),
		Role:   req.Role,
		UserId: id,
		Filter: &pb.Filter{
			UserId:  0,
			GroupId: uint32(groupId),
		},
	}

	listResp, err := service.FeedClient.List(context.Background(), listReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
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
