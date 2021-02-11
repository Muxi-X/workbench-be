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

// GetFileFolderInfoList ... 获取文件夹列表
func GetFileFolderInfoList(c *gin.Context) {
	log.Info("project getFileFolderInfoList function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获得请求
	var req GetFileInfoListRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	resp, err := service.ProjectClient.GetFileFolderInfoList(context.Background(), &pbp.GetInfoByIdsRequest{
		List: req.Ids,
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
