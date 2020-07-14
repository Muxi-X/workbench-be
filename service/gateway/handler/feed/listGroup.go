package feed

import (
	"context"
	//"fmt"
	"log"
	"strconv"

	"muxi-workbench/pkg/handler"
	//"muxi-workbench-feed-client/tracer"
	pb "muxi-workbench-feed/proto"
	"muxi-workbench/pkg/errno"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

// Feed 的 ListGroup 接口
func ListGroup(c *gin.Context) {
	log.Info("Feed list Group function called.")
	// 从 Query Param 中获得 limit 和 lastid
	var limit int
	var lastid int
	var err error
	var gid int

	// 多一个获取 gid 的步骤
	gid, err = strconv.Atoi(c.Param("gid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	limit, err = strconv.Atoi(c.DefautlQuery("limit", "50"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	lastid, err = strconv.Atoi(c.Query("lastid", "0"))
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
		LastId: lastid,
		Limit:  limit,
		Role:   req.Role,
		UserId: req.Userid,
		Filter: &pb.Filter{
			UserId:  0,
			GroupId: gid,
		},
	}

	listResp, err2 := FeedClient.List(context.Background(), listReq)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 构造返回 response
	var resp listResponse
	for i := 0; i < len(resp.FeedItem); i++ {
		resp.FeedItem = append(resp.FeedItem, feedItem{
			Id:          listResp.FeedItem[i].Id,
			Action:      listResp.FeedItem[i].Actionhow,
			ShowDivider: listResp.FeedItem[i].ShowDivider,
			Date:        listResp.FeedItem[i].Date,
			Time:        listResp.FeedItem[i].Time,
			User:        listResp.FeedItem[i].User,
			Source:      listResp.FeedItem[i].Source,
		})
	}
	resp.Count = listResp.Count

	SendResponse(c, nil, resp)
}
