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

// 只用调用一次 get member
// 不用从 token 获取 userid
func GetMembers(c *gin.Context) {
	log.Info("Project get member function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 要从 query param 获取 lastid limit page pagination
	var pid int
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

	var pageBool bool

	if pagination == 1 {
		pageBool = true
	}

	// 构造请求
	getMemReq := &pbp.GetMemberListRequest{
		ProjectId:  uint32(pid),
		Lastid:     uint32(lastid),
		Offset:     uint32(page),
		Limit:      uint32(limit),
		Pagination: pageBool,
	}

	getMemResp, err2 := service.ProjectClient.GetMembers(context.Background(), getMemReq)
	if err2 != nil {
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
			GroupName: getMemResp.List[i].GroupName,
			Role:      getMemResp.List[i].Role,
		})
	}

	SendResponse(c, nil, resp)
}
