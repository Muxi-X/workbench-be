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

// GetMemberList 获取组别内成员列表
func GetMemberList(c *gin.Context) {
	log.Info("Member List function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var limit int
	var groupID int
	var page int
	var err error
	pagination := true

	groupID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	page, err = strconv.Atoi(c.Query("page"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
	if page == -1 {
		pagination = false
	}

	// 构造 MemberList 请求
	MemberListReq := &tpb.MemberListRequest{
		GroupId:    uint32(groupID),
		Offset:     uint32(limit * page),
		Limit:      uint32(limit),
		Pagination: pagination,
	}

	// 向 GetMemberList 服务发送请求
	listResp, err := service.TeamClient.GetMemberList(context.Background(), MemberListReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	var resp memberListResponse
	for i := 0; i < len(listResp.List); i++ {
		resp.Members = append(resp.Members, member{
			ID:        listResp.List[i].Id,
			Name:      listResp.List[i].Name,
			TeamID:    listResp.List[i].TeamId,
			GroupID:   listResp.List[i].GroupId,
			GroupName: listResp.List[i].GroupName,
			Email:     listResp.List[i].Email,
			Avatar:    listResp.List[i].Avatar,
		})
	}
	resp.Count = listResp.Count

	SendResponse(c, nil, resp)
}
