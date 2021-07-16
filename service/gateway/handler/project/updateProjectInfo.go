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

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UpdateProjectInfo updates a project's info
// @Summary update project info api
// @Description 修改项目详情
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body ProjectUpdateRequest true "update_request"
// @Param project_id query int true "项目 id"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /project [put]
func UpdateProjectInfo(c *gin.Context) {
	log.Info("Project updateProjectInfo function call", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 projectID
	projectID := c.MustGet("projectID").(uint32)

	// 获取请求
	var req ProjectUpdateRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userid
	userID := c.MustGet("userID").(uint32)

	// 构造 update 请求
	updateReq := &pbp.UpdateProjectInfoRequest{
		Id:    uint32(projectID),
		Name:  req.ProjectName,
		Intro: req.Intro,
	}

	// 发送 update 请求
	_, err := service.ProjectClient.UpdateProjectInfo(context.Background(), updateReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	/* --- 新增 feed 动态 --- */

	// 构造 push 请求
	pushReq := &pbf.PushRequest{
		Action: "编辑",
		UserId: userID,
		Source: &pbf.Source{
			Kind:        2,
			Id:          uint32(projectID),
			Name:        "",
			ProjectId:   uint32(projectID),
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
