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
func GetDocDetail(c *gin.Context) {
	log.Info("project getDocDetail function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	docID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	getDocDetailResp, err := service.ProjectClient.GetDocDetail(context.Background(), &pbp.GetRequest{
		Id: uint32(docID),
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
