package user

import (
	"context"
	"strconv"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pb "muxi-workbench-user/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 暂时不知道 router
// List 通过 group 和 team 获取 userlist
// 通过 param 获取 page last_id
func List(c *gin.Context) {
	log.Info("User getInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 limit
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 page
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 group 和 team
	var req listRequest
	if err := c.BindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造请求给 list
	listReq := &pb.ListRequest{
		LastId: 0,
		Offset: uint32(page * limit),
		Limit:  uint32(limit),
		Team:   req.Team,
		Group:  req.Group,
	}

	// 发送请求
	listResp, err := service.UserClient.List(context.Background(), listReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	var resp = &listResponse{Count: listResp.Count}
	for _, item := range listResp.List {
		resp.List = append(resp.List, user{
			Id:     item.Id,
			Nick:   item.Nick,
			Name:   item.Name,
			Avatar: item.Avatar,
			Email:  item.Email,
			Role:   item.Role,
			Team:   item.Team,
			Group:  item.Group,
		})
	}

	SendResponse(c, nil, resp)
}
