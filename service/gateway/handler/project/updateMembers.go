package project

import (
	"context"

	pbf "muxi-workbench-feed/proto"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// UpdateMembers updates the members in the project
// @Summary update project member api
// @Description 修改项目成员
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body UpdateMemberRequest true "update_member_request"
// @Param project_id query int true "项目 id"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /project/member [put]
func UpdateMembers(c *gin.Context) {
	log.Info("Project member update function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 projectID
	projectID := c.MustGet("projectID").(uint32)

	// 获取请求
	var req UpdateMemberRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userID
	userID := c.MustGet("userID").(uint32)

	// 构造请求
	// 这里 list 应该是 uint32 表示 uid
	updateMemReq := &pbp.UpdateMemberRequest{
		Id: projectID,
	}

	for i := 0; i < len(req.UserList); i++ {
		updateMemReq.List = append(updateMemReq.List, req.UserList[i])
	}

	// 发送请求
	_, err := service.ProjectClient.UpdateMembers(context.Background(), updateMemReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	/* --- 新增 feed 动态 --- */

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "加入",
		UserId: userID,
		Source: &pbf.Source{
			Kind:        2,
			Id:          projectID,
			Name:        "",
			ProjectId:   projectID,
			ProjectName: "",
		},
	}

	// 向 feed 发送请求
	_, err = service.FeedClient.Push(context.Background(), pushReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
