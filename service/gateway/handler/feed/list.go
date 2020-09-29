package feed

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	pb "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"

	"github.com/gin-gonic/gin"
)

// Feed 的 List 接口
// 需要从 token 获取 userid ?
func List(c *gin.Context) {
	log.Info("feed List function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从 Query Param 中获得 limit 和 lastid
	var limit int
	var lastId int
	var err error

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "50"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	lastId, err = strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	var req listRequest
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
			GroupId: 0,
		},
	}

	listResp, err := service.FeedClient.List(context.Background(), listReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	var resp listResponse
	for i := 0; i < len(resp.FeedItem); i++ {
		resp.FeedItem = append(resp.FeedItem, feedItem{
			Id:          listResp.List[i].Id,
			Action:      listResp.List[i].Action,
			ShowDivider: listResp.List[i].ShowDivider,
			Date:        listResp.List[i].Date,
			Time:        listResp.List[i].Time,
			User: user{
				Name:      listResp.List[i].User.Name,
				Id:        listResp.List[i].User.Id,
				AvatarUrl: listResp.List[i].User.AvatarUrl,
			},
			Source: source{
				Kind:        listResp.List[i].Source.Kind,
				Id:          listResp.List[i].Source.Id,
				Name:        listResp.List[i].Source.Name,
				ProjectId:   listResp.List[i].Source.ProjectId,
				ProjectName: listResp.List[i].Source.ProjectName,
			},
		})
	}
	resp.Count = listResp.Count

	SendResponse(c, nil, resp)
}
