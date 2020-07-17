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

// 需要调用 get 和 listcomment
func Get(c *gin.Context) {
	log.Info("Status get function call.")

	// 从 Query Param 中获取 lastid 和 limit
	var limit int
	var lastid int
	var sid int
	var page int
	var err error

	sid, err = strconv.Atoi(c.Param("sid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

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

	// 构造 get status 请求并发送
	getReq := &pbs.GetRequest{
		Id: sid,
	}

	getResp, err2 := StatusClient.Get(context.Background(), &getReq)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err2.Error())
		return
	}

	// 构造 listcomment 请求并发送
	listComReq := &pbs.CommentListRequest{
		StatusId: sid,
		Offset:   page,
		Limit:    limit,
		Lastid:   lastid,
	}

	listComResp, err3 := StatusClient.ListComment(context.Background(), listComReq)
	if err3 != nil {
		log.Fatalf("Could not greet: %v", err3)
		SendError(c, errno.InternalServerError, nil, err3.Error())
		return
	}

	// 构造返回 response
	resp := getResponse{
		Sid:      getResp.Id,
		Title:    getResp.Title,
		Content:  getResp.Content,
		UserId:   getResp.UserId,
		Time:     getResp.Time,
		Avatar:   getResp.Avatar,
		Username: getResp.Username,
		Count:    listComResp.Count,
	}

	for i := 0; i < len(listComResp.Comment); i++ {
		resp.Commentlist = append(resp.Commentlist, comment{
			Cid:      listComResp.List[i].Id,
			Uid:      listComResp.List[i].Userid,
			Username: listComResp.List[i].Username,
			Avatar:   listComResp.List[i].Avatar,
			Time:     listComResp.List[i].Time,
			Content:  listComResp.List[i].Content,
		})
	}

	// 返回 response
	SendResponse(c, nil, resp)
}
