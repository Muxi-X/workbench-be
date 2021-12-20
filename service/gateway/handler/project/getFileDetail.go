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

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// GetFileDetail gets a file's detail
// @Summary get file detail api
// @Description 获取某个文件的详情
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param file_id path int true "file_id"
// @Param id path int true "father_id"
// @Param project_id query int true "project_id"
// @Success 200 {object} GetFileDetailResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /file/file/{id}/children/{file_id} [get]
func GetFileDetail(c *gin.Context) {
	log.Info("project getFileDetail function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	projectID := c.MustGet("projectID").(uint32)
	fileID, err := strconv.Atoi(c.Param("file_id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	fatherID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	getFileDetailResp, err := service.ProjectClient.GetFileDetail(context.Background(), &pbp.GetFileDetailRequest{
		ProjectId: projectID,
		Id:        uint32(fileID),
		FatherId:  uint32(fatherID),
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回结果
	resp := GetFileDetailResponse{
		Id:         getFileDetailResp.Id,
		Url:        getFileDetailResp.Url,
		Creator:    getFileDetailResp.Creator,
		CreateTime: getFileDetailResp.CreateTime,
	}

	SendResponse(c, nil, resp)
}
