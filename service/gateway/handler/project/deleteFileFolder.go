package project

import (
	"context"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// DeleteFileFolder ... 删除文件夹
func DeleteFileFolder(c *gin.Context) {
	log.Info("project deleteFileFolder function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	folderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	deleteDocFolderReq := &pbp.GetRequest{
		Id: uint32(folderId),
	}

	_, err = service.ProjectClient.DeleteFileFolder(context.Background(), deleteDocFolderReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
