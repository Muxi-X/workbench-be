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

// GetDocDetail gets a doc's detail
// @Summary get doc detail api
// @Description 获取某个文档的详情
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param file_id path int true "doc_id"
// @Param id path int true "fahter_id"
// @Param project_id query int true "project_id"
// @Success 200 {object} handler.Response{data=GetDocDetailResponse}
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /file/doc/{id}/children/{file_id} [get]
func GetDocDetail(c *gin.Context) {
	log.Info("project getDocDetail function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	docID, err := strconv.Atoi(c.Param("file_id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}
	fatherID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	getDocDetailResp, err := service.ProjectClient.GetDocDetail(context.Background(), &pbp.GetFileDetailRequest{
		Id:       uint32(docID),
		FatherId: uint32(fatherID),
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回结果
	resp := GetDocDetailResponse{
		Id:           getDocDetailResp.Id,
		Title:        getDocDetailResp.Title,
		Content:      getDocDetailResp.Content,
		Creator:      getDocDetailResp.Creator,
		CreateTime:   getDocDetailResp.CreateTime,
		LastEditor:   getDocDetailResp.LastEditor,
		LastEditTime: getDocDetailResp.LastEditTime,
	}

	SendResponse(c, nil, resp)
}
