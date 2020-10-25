package project

import (
	"context"
	"strconv"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetProjectInfo gets a project's information by its id
func GetProjectInfo(c *gin.Context) {
	log.Info("project getProjectInfo function call", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 projectID
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 发送请求
	getProInfoResp, err := service.ProjectClient.GetProjectInfo(context.Background(), &pbp.GetRequest{
		Id: uint32(projectID),
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	resp := GetProjectInfoResponse{
		ProjectID:   getProInfoResp.Id,
		ProjectName: getProInfoResp.Name,
		Intro:       getProInfoResp.Intro,
		UserCount:   getProInfoResp.UserCount,
	}

	// 返回结果
	SendResponse(c, nil, resp)
}
