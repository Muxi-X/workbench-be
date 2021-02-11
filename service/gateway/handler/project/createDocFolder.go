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

// CreateDocFolder ... 建立文档夹
func CreateDocFolder(c *gin.Context) {
	log.Info("project createDocFolder function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req CreateFolderRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	userID := c.MustGet("userID").(uint32)

	createDocFolderReq := &pbp.CreateFolderRequest{
		Name:       req.Name,
		CreatorId:  userID,
		ProjectId:  req.ProjectId,
		FatherId:   req.FatherId,
		FatherType: req.FatherType,
	}

	_, err := service.ProjectClient.CreateDocFolder(context.Background(), createDocFolderReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
