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
	"muxi-workbench-project/model"
	pbp "muxi-workbench-project/proto"
	"muxi-workbench/pkg/constvar"
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
	for i, file := range fileList {
		id, _ := strconv.Atoi(file.Id)
		// 发送请求获取name
		if file.Type { // file folder
			isDeleted, err := model.AdjustSelfIfExist(uint32(id), constvar.FileFolderCode)
			if err != nil {
				SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
				return
			}
			if isDeleted { // 存在 redis 返回 1, 说明被删
				fileList = append(fileList[:i], fileList[i+1:]...) // 删除中间1个元素
			}
			getFileNameResp, err := service.ProjectClient.GetFileOrDocName(context.Background(), &pbp.GetFileOrDocNameRequest{
				Id:   uint32(id),
				Type: uint32(constvar.FileFolderCode),
			})
			if err != nil {
				SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
				return
			}
			file.Name = getFileNameResp.Name
		} else {
			isDeleted, err := model.AdjustSelfIfExist(uint32(id), constvar.FileCode)
			if err != nil {
				SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
				return
			}
			if isDeleted { // 存在 redis 返回 1, 说明被删
				fileList = append(fileList[:i], fileList[i+1:]...)
			}
			getFileNameResp, err := service.ProjectClient.GetFileOrDocName(context.Background(), &pbp.GetFileOrDocNameRequest{
				Id:   uint32(id),
				Type: uint32(constvar.FileCode),
			})
			if err != nil {
				SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
				return
			}
			file.Name = getFileNameResp.Name
		}
	}

	docList := FormatChildren(getProInfoResp.DocChildren)
	for i, doc := range docList {
		// 发送请求获取name
		id, _ := strconv.Atoi(doc.Id)
		if doc.Type { // doc folder
			isDeleted, err := model.AdjustSelfIfExist(uint32(id), constvar.DocFolderCode)
			if err != nil {
				SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
				return
			}
			if isDeleted { // 存在 redis 返回 1, 说明被删
				fileList = append(fileList[:i], fileList[i+1:]...) // 删除中间1个元素
			}
			getDocNameResp, err := service.ProjectClient.GetFileOrDocName(context.Background(), &pbp.GetFileOrDocNameRequest{
				Id:   uint32(id),
				Type: uint32(constvar.DocFolderCode),
			})
			if err != nil {
				SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
				return
			}
			doc.Name = getDocNameResp.Name
		} else {
			isDeleted, err := model.AdjustSelfIfExist(uint32(id), constvar.DocCode)
			if err != nil {
				SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
				return
			}
			if isDeleted { // 存在 redis 返回 1, 说明被删
				fileList = append(fileList[:i], fileList[i+1:]...) // 删除中间1个元素
			}
			getDocNameResp, err := service.ProjectClient.GetFileOrDocName(context.Background(), &pbp.GetFileOrDocNameRequest{
				Id:   uint32(id),
				Type: uint32(constvar.DocCode),
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
