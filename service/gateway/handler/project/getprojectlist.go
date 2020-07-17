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

// 需要调用 list
func GetProjectList(c *gin.Context) {
	log.Info("Project list get function call.")

	// 获取 lastid page limit pagination
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

	// 从请求获取 userid
	var req getProjectListRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 构造请求
	getProListReq := pbp.GetProjectListRequest{
		UserId:     req.UserId,
		Lastid:     lastid,
		Offset:     page,
		Limit:      limit,
		Pagination: pagination,
	}

	getProListResp, err2 := ProjectClient.GetProjectList(context.Background(), getProListReq)
	if err2 != nil {
		log.Fatalf("Could not greet: %v", err2)
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 构造response
	var resp getProjectListRequest
	for i := 0; i < len(getProListResp.List); i++ {
		resp.Projectlist = append(resp.Projectlist, projectListItem{
			Id:   getProListResp.List[i].Id,
			Name: getProListResp.List[i].Name,
			Logo: getProListResp.List[i].Logo,
		})
	}

	SendResponse(c, nil, resp)

}
