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

// UpdateFilePosition ... 移动文件
func UpdateFilePosition(c *gin.Context) {
	log.Info("Project file position update function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 fatherId 和 oldFatherId
	oldFatherId, err := strconv.Atoi(c.Param("father_id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取请求
	var req UpdateFilePositionRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造请求
	// 这里 list 应该是 uint32 表示 uid
	updateFilePositionReq := &pbp.UpdateFilePositionRequest{
		FileId:                uint32(id),
		OldFatherId:           uint32(oldFatherId),
		FatherId:              req.FatherId,
		FatherType:            req.FatherType,
		Type:                  uint32(req.Type),
		ChildrenPositionIndex: req.ChildrenPositionIndex,
	}

	// 发送请求
	_, err = service.ProjectClient.UpdateFilePosition(context.Background(), updateFilePositionReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
