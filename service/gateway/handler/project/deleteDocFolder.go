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

// DeleteDocFolder ... 删除文档夹
func DeleteDocFolder(c *gin.Context) {
	log.Info("project deleteDocFolder function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	folderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取请求
	var req DeleteFolderRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 userID
	userID := c.MustGet("userID").(uint32)
	role := c.MustGet("role").(uint32)

	deleteDocFolderReq := &pbp.DeleteRequest{
		Id:         uint32(folderId),
		FatherId:   req.FatherId,
		FatherType: req.FatherType,
		UserId:     userID,
		Role:       role,
	}

	_, err = service.ProjectClient.DeleteDocFolder(context.Background(), deleteDocFolderReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
