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

// CreateProject creates new project
// @Summary create a project api
// @Description 新建项目
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body CreateProjectRequest true "create_project_request"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /project [post]
func CreateProject(c *gin.Context) {
	log.Info("project createProject function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req CreateProjectRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	userID := c.MustGet("userID").(uint32)
	teamID := c.MustGet("teamID").(uint32)

	createProjectReq := &pbp.CreateProjectRequest{
		Name:      req.Name,
		Intro:     req.Intro,
		TeamId:    teamID,
		CreatorId: userID,
	}

	resp, err := service.ProjectClient.CreateProject(context.Background(), createProjectReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
	}

	pushReq := &pbf.PushRequest{
		Action: "创建",
		UserId: userID,
		Source: &pbf.Source{
			Kind:        2,
			Id:          uint32(0),
			Name:        "",
			ProjectId:   resp.Id,
			ProjectName: req.Name,
		},
	}

	// 向 feed 发送请求
	_, err = service.FeedClient.Push(context.Background(), pushReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, nil)
}
