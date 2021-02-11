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

// UpdateFileFolder ... 修改文件夹
func UpdateFileFolder(c *gin.Context) {
	log.Info("Project updateFileFolder function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 folderID 和请求
	var err error
	var folderID int

	folderID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
	}

	// 获取请求体
	var req UpdateFolderRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 updateReq 并发送请求
	updateReq := &pbp.UpdateFolderRequest{
		FolderId: uint32(folderID),
		Name:     req.Name,
	}

	_, err = service.ProjectClient.UpdateFileFolder(context.Background(), updateReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
