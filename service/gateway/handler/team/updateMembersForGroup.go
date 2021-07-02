package team

import (
	"context"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	tpb "muxi-workbench-team/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UpdateMembersForGroup ... 更新组别内成员
// @Summary update members of group api
// @Description 更新组别成员
// @Tags group
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body UpdateMembersRequest true "update_members_request"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /team/group/members [put]
func UpdateMembersForGroup(c *gin.Context) {
	log.Info("Members update in Group function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req UpdateMembersRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 updateMembers 请求
	updateMembersReq := &tpb.UpdateMembersRequest{
		GroupId:  req.GroupID,
		UserList: req.UserList,
	}

	// 向 UpdateMembersForGroup 服务发送请求
	_, err := service.TeamClient.UpdateMembersForGroup(context.Background(), updateMembersReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
