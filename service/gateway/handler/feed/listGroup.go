package feed

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	pb "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/pkg/token"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"

	"github.com/gin-gonic/gin"
)

// Feed 的 ListGroup 接口
// 不需要从 token 获取 userid
func ListGroup(c *gin.Context) {
	log.Info("Feed listGroup function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	// 从 Query Param 中获得 limit 和 lastid
	var limit int
	var lastid int
	var err error
	var gid int

	// 多一个获取 gid 的步骤
	gid, err = strconv.Atoi(c.Param("gid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		log.Fatal("feed listGroup, get param:gid fatal",
			zap.String("reason", err.Error()))
		return
	}

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "50"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		log.Fatal("feed listGroup, get param:limit fatal",
			zap.String("reason", err.Error()))
		return
	}

	lastid, err = strconv.Atoi(c.DefaultQuery("lastid", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		log.Fatal("feed listGroup, get param:lastid fatal",
			zap.String("reason", err.Error()))
		return
	}

	var req listRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		log.Fatal("feed listGroup, bind request fatal",
			zap.String("reason", err.Error()))
		return
	}

	// 获取 userid
	raw, ifexists := c.Get("context")
	if !ifexists {
		SendBadRequest(c, errno.ErrTokenInvalid, nil, "Context not exists")
		log.Fatal("feed listGroup, get userid raw from context fatal",
			zap.String("reason", "maybe raw in the context not exists"))
		return
	}
	ctx, ok := raw.(*token.Context)
	if !ok {
		SendError(c, errno.ErrValidation, nil, "Context assign failed")
		log.Fatal("feed listGroup, get userid from raw fatal",
			zap.String("reason", "maybe raw type assertion fatal"))
		return
	}

	listReq := &pb.ListRequest{
		LastId: uint32(lastid),
		Limit:  uint32(limit),
		Role:   req.Role,
		UserId: uint32(ctx.ID),
		Filter: &pb.Filter{
			UserId:  0,
			GroupId: uint32(gid),
		},
	}

	listResp, err2 := service.FeedClient.List(context.Background(), listReq)
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err.Error())
		log.Fatal("feed listGroup, get response from server fatal",
			zap.String("reason", err2.Error()))
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
