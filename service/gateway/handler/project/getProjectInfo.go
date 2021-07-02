package project

import (
	"context"

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
// @Summary get project info api
// @Description 获取项目的详情
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param project_id query int true "project_id"
// @Success 200 {object} GetProjectInfoResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /project [get]
func GetProjectInfo(c *gin.Context) {
	log.Info("project getProjectInfo function call", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 projectID
	projectID := c.MustGet("projectID").(uint32)

	// 发送请求
	getProInfoResp, err := service.ProjectClient.GetProjectInfo(context.Background(), &pbp.GetRequest{
		Id: projectID,
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 解析结果
	docList := FormatChildren(getProInfoResp.DocChildren)
	fileList := FormatChildren(getProInfoResp.FileChildren)

	// 构造返回 response
	resp := GetProjectInfoResponse{
		ProjectID:    getProInfoResp.Id,
		ProjectName:  getProInfoResp.Name,
		Intro:        getProInfoResp.Intro,
		UserCount:    getProInfoResp.UserCount,
		DocChildren:  docList,
		FileChildren: fileList,
	}

	// 返回结果
	SendResponse(c, nil, resp)
}
