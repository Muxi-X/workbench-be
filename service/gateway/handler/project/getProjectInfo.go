package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"
	"strconv"
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
	fileList := FormatChildren(getProInfoResp.FileChildren)
	for _, file := range fileList {
		// 发送请求获取name
		id, _ := strconv.Atoi(file.Id)
		if file.Type { // file folder
			getFileNameResp, err := service.ProjectClient.GetFileOrDocName(context.Background(), &pbp.GetFileOrDocNameRequest{
				Id:   uint32(id),
				Type: 3,
			})
			if err != nil {
				SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
				return
			}
			file.Name = getFileNameResp.Name
		} else {
			getFileNameResp, err := service.ProjectClient.GetFileOrDocName(context.Background(), &pbp.GetFileOrDocNameRequest{
				Id:   uint32(id),
				Type: 1,
			})
			if err != nil {
				SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
				return
			}
			file.Name = getFileNameResp.Name
		}
	}

	docList := FormatChildren(getProInfoResp.DocChildren)
	for _, doc := range docList {
		// 发送请求获取name
		id, _ := strconv.Atoi(doc.Id)
		if doc.Type { // doc folder
			getDocNameResp, err := service.ProjectClient.GetFileOrDocName(context.Background(), &pbp.GetFileOrDocNameRequest{
				Id:   uint32(id),
				Type: 4,
			})
			if err != nil {
				SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
				return
			}
			doc.Name = getDocNameResp.Name
		} else {
			getDocNameResp, err := service.ProjectClient.GetFileOrDocName(context.Background(), &pbp.GetFileOrDocNameRequest{
				Id:   uint32(id),
				Type: 2,
			})
			if err != nil {
				SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
				return
			}
			doc.Name = getDocNameResp.Name
		}
	}

	// 构造返回 response
	resp := GetProjectInfoResponse{
		ProjectID:    getProInfoResp.Id,
		ProjectName:  getProInfoResp.Name,
		Intro:        getProInfoResp.Intro,
		UserCount:    getProInfoResp.UserCount,
		DocChildren:  docList,
		FileChildren: fileList,
		Time:         getProInfoResp.Time,
		CreatorName:  getProInfoResp.CreatorName,
	}

	// 返回结果
	SendResponse(c, nil, resp)
}
