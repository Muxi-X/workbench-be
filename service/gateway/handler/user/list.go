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

// List ... 获取 userlist
// @Summary get user_list api
// @Description 通过 group 和 team 获取 userlist
// @Tags user
// @Accept  application/json
// @Produce  application/json
// @Param limit query string false "limit"
// @Param page query string false "page"
// @Param Authorization header string false "token 用户令牌"
// @Param object body ListRequest  false "get_user_list_request"
// @Security ApiKeyAuth
// @Success 200 {object} ListResponse
// @Router /user/list [get]
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
	var req ListRequest
	if err := c.BindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	lastId, err := strconv.Atoi(c.DefaultQuery("lastid", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造请求给 list
	listReq := &pb.ListRequest{
		LastId: uint32(lastId),
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
	var resp = &ListResponse{Count: listResp.Count}
	for _, item := range listResp.List {
		resp.List = append(resp.List, User{
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
