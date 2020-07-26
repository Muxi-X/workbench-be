package project

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
)

// 需要调用 list
func GetProjectList(c *gin.Context) {
	log.Info("Project list get function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

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

	var pageBool bool
	if pagination == 1 {
		pageBool = true
	}

	// 构造请求
	getProListReq := &pbp.GetProjectListRequest{
		UserId:     req.UserId,
		Lastid:     uint32(lastid),
		Offset:     uint32(page),
		Limit:      uint32(limit),
		Pagination: pageBool,
	}

	getProListResp, err2 := service.ProjectClient.GetProjectList(context.Background(), getProListReq)
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	// 构造response
	var resp getProjectListResponse
	for i := 0; i < len(getProListResp.List); i++ {
		resp.Projectlist = append(resp.Projectlist, projectListItem{
			Id:   getProListResp.List[i].Id,
			Name: getProListResp.List[i].Name,
			Logo: getProListResp.List[i].Logo,
		})
	}

	SendResponse(c, nil, resp)

}
