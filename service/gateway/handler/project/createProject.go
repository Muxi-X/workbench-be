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
func CreateProject(c *gin.Context) {
	log.Info("project createProject function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req CreateProjectRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	userID := c.MustGet("userID").(uint32)

	createProjectReq := &pbp.CreateProjectRequest{
		Name:   req.Name,
		Intro:  req.Intro,
		TeamId: req.TeamId,
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
