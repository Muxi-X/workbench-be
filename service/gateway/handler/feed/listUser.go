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

// Feed 的 ListUser 接口
func ListUser(c *gin.Context) {
	log.Info("Feed listUser function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	// 从 Query Param 中获得 limit 和 lastid
	var limit int
	var lastid int
	var err error
	var uid int

	// 多一个获取 uid 的步骤
	uid, err = strconv.Atoi(c.Param("uid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "50"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	lastid, err = strconv.Atoi(c.DefaultQuery("lastid", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	var req listRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	listReq := &pb.ListRequest{
		LastId: uint32(lastid),
		Limit:  uint32(limit),
		Role:   req.Role,
		UserId: req.Userid,
		Filter: &pb.Filter{
			UserId:  uint32(uid),
			GroupId: 0,
		},
	}

	listResp, err2 := service.FeedClient.List(context.Background(), listReq)
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err.Error())
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
