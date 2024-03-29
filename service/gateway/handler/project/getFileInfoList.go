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

// GetFileInfoList ... 获取文件列表
// @Summary get file info list api
// @Description 获取文件列表
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param ids query []int true "folder_ids 是一个数组"
// @Param project_id query int true "project_id"
// @Success 200 {object} GetFileInfoListResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /file/files [get]
func GetFileInfoList(c *gin.Context) {
	log.Info("project getFileInfoList function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从 query 获得 ids
	rawIds, isvalid := c.GetQueryArray("ids")
	if !isvalid {
		SendBadRequest(c, errno.ErrQuery, nil, "no query parameters", GetLine())
		return
	}

	// 转换
	var ids []uint32
	for _, v := range rawIds {
		id, err := strconv.Atoi(v)
		if err != nil {
			SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
			return
		}
		ids = append(ids, uint32(id))
	}
	projectID := c.MustGet("projectID").(uint32)

	resp, err := service.ProjectClient.GetFileInfoList(context.Background(), &pbp.GetInfoByIdsRequest{
		ProjectId: projectID,
		List:      ids,
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	var list []*FileInfoItem
	for _, v := range resp.List {
		list = append(list, &FileInfoItem{
			Id:   v.Id,
			Name: v.Title,
		})
	}

	SendResponse(c, errno.OK, &GetFileInfoListResponse{
		Count: uint32(len(list)),
		List:  list,
	})
}
