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

// GetDocFolderInfoList ... 获取文档夹列表
func GetDocFolderInfoList(c *gin.Context) {
	log.Info("project getDocFolderInfoList function call.",
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

	resp, err := service.ProjectClient.GetDocFolderInfoList(context.Background(), &pbp.GetInfoByIdsRequest{
		List: ids,
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	var list []*FileInfoItem
	for _, v := range resp.List {
		list = append(list, &FileInfoItem{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	SendResponse(c, errno.OK, &GetFileInfoListResponse{
		Count: uint32(len(list)),
		List:  list,
	})
}
