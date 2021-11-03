package team

import (
	"context"
	"strconv"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	tpb "muxi-workbench-team/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetMemberList ... 获取组别内成员列表
// @Summary list members api
// @Description 根据 groupID 拉取 members 列表 (0 -> all)
// @Tags group
// @Accept  application/json
// @Produce  application/json
// @Param id path int true "group_id"
// @Param Authorization header string true "token 用户令牌"
// @Param limit query int false "limit"
// @Param page query int false "page 从 0 开始计数， 如果传入非负数或者不传值则不分页"
// @Success 200 {object} MemberListResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /team/group/members/{id} [get]
func GetMemberList(c *gin.Context) {
	log.Info("Members List function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	pagination := true

	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "-1"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}
	if page < 0 {
		pagination = false
	}

	// 构造 memberList 请求
	memberListReq := &tpb.MemberListRequest{
		GroupId:    uint32(groupID),
		Offset:     uint32(limit * (page)),
		Limit:      uint32(limit),
		Pagination: pagination,
	}

	// 向 GetMemberList 服务发送请求
	listResp, err := service.TeamClient.GetMemberList(context.Background(), memberListReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	var resp MemberListResponse
	for _, item := range listResp.List {
		resp.Members = append(resp.Members, Member{
			ID:        item.Id,
			Name:      item.Name,
			TeamID:    item.TeamId,
			GroupID:   item.GroupId,
			GroupName: item.GroupName,
			Email:     item.Email,
			Avatar:    item.Avatar,
			Role:      item.Role,
		})
	}
	resp.Count = listResp.Count

	SendResponse(c, nil, resp)
}
