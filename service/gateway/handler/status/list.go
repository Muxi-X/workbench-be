package status

import (
	"context"
	//"fmt"
	"log"

	//tracer "muxi-workbench-status-client/tracer"
	pbs "muxi-workbench-status/proto"
	handler "muxi-workbench/pkg/handler"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
)

// 只用调用一次 list  lastid limit page 要从 query param 获取
func List(c *gin.Context) {
	log.Info("Status list function call")

	// 获取 gid 和 limt lastid
	var limt int
	var lastid int
	var err error

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	lastid, err = strconv.Atoi(c.DefaultQuery("lastid", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 获取 page
	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 构造 list 请求
	listReq := &pbs.ListRequest{
		Lastid: lastid,
		Offset: page,
		Limit:  limit,
		Group:  0,
		Uid:    0,
	}

	listResp, err2 := StatusClient.List(context.Background(), listReq)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 构造返回 response
	var resp listResponse
	for i := 0; i < len(listResp.Status); i++ {
		resp.Status = append(resp.Status, status{
			Id:       listResp.Status[i].Id,
			Title:    listResp.Status[i].Title,
			Content:  listResp.Status[i].Content,
			UserId:   listResp.Status[i].UserId,
			Time:     listResp.Status[i].Time,
			Avatar:   listResp.Status[i].Avatar,
			Username: listResp.Status[i].Username,
		})
	}
	resp.Count = listResp.Count

	SendResponse(c, nil, resp)
}
