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

// 只用调用一次 list  lastid limit 要从 query param 获取 还要获取gid
func List(c *gin.Context) {
	log.Info("Status list group function call")

	// 获取 gid 和 limt lastid
	var limt int
	var lastid int
	var gid int
	var page int
	var err error

	gid, err = strconv.Atoi(c.Param("gid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	lastid, err = strconv.Atoi(c.Query("lastid", "0"))
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
		Group:  gid,
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
			Id:       listResp.List[i].Id,
			Title:    listResp.List[i].Title,
			Content:  listResp.List[i].Content,
			UserId:   listResp.List[i].UserId,
			Time:     listResp.List[i].Time,
			Avatar:   listResp.List[i].Avatar,
			Username: listResp.List[i].Username,
		})
	}
	resp.Count = listResp.Count

	SendResponse(c, nil, resp)
}
