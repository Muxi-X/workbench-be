package user

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	pb "muxi-workbench-user/proto"
	// "muxi-workbench-gateway/pkg/token"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"

	"github.com/gin-gonic/gin"
)

// 暂时不知道 router
// List 通过 group 和 team 获取 userlist
// 通过 param 获取 limit offset lastid
func List(c *gin.Context) {
	log.Info("User getInfo function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 offset/page 和 limt lastid
	var limit int
	var lastid int
	var page int
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

	// 从前端获取 group 和 team
	var req listRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 构造请求给 list
	listReq := &pb.ListRequest{
		LastId: uint32(lastid),
		Offset: uint32(page),
		Limit:  uint32(limit),
		Team:   req.Team,
		Group:  req.Group,
	}

	// 发送请求
	listResp, err2 := service.UserClient.List(context.Background(), listReq)
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err2.Error())
		return
	}

	// 构造返回 response
	var resp listResponse
	for i := 0; i < len(listResp.List); i++ {
		resp.List = append(resp.List, user{
			Id:     listResp.List[i].Id,
			Nick:   listResp.List[i].Nick,
			Name:   listResp.List[i].Name,
			Avatar: listResp.List[i].Avatar,
			Email:  listResp.List[i].Email,
			Role:   listResp.List[i].Role,
			Team:   listResp.List[i].Team,
			Group:  listResp.List[i].Group,
		})
	}
	resp.Count = listResp.Count

	SendResponse(c, nil, resp)
}
