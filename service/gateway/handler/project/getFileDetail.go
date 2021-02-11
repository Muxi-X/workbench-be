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
func GetFileDetail(c *gin.Context) {
	log.Info("project getFileDetail function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	fileID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	getFileDetailResp, err := service.ProjectClient.GetFileDetail(context.Background(), &pbp.GetRequest{
		Id: uint32(fileID),
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
