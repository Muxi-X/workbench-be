package handler

import (
	"context"
	//"fmt"
	"log"

	//tracer "muxi-workbench-project-client/tracer"
	pbp "muxi-workbench-project/proto"
	handler "muxi-workbench/pkg/handler"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro"
	opentracingWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

// 只用调用一次 get member
func GetMembers(c *gin.Context) {
	log.Info("Project get member function call.")

	// 要从 query param 获取 lastid limit page pagination
	var limit int
	var lastid int
	var page int
	var pagination int
	var err error

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	lastid, err = strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	page, err = strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	pagination, err = strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 获取 pid
	pid, err = strconv.Atoi(c.Param("pid"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 构造请求
	getMemReq := &pbp.GetMemberListRequest{
		ProjectId:  pid,
		Lastid:     lastid,
		Offset:     page,
		Limit:      limit,
		Pagination: pagination,
	}

	getMemResp, err2 := ProjectClient.GetMembers(context.Background(), getMemReq)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 构造返回 resp 错了，要改
	resp := getMemberResponse{
		Count: getMemResp.Count,
	}
	for i := 0; i < len(getMemResp.List); i++ {
		resp.Memberlist = append(resp.Memberlist, memberListItem{
			Id:        getMemResp.List[i].Id,
			Name:      getMemResp.List[i].Name,
			Avatar:    getMemResp.List[i].Avatar,
			Groupname: getMemResp.List[i].GroupName,
			Role:      getMemResp.List[i].Role,
		})
	}

	SendResponse(c, nil, resp)
}
