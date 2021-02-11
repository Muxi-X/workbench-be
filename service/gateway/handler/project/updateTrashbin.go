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

// UpdateTrashbin ... 用于恢复资源
// 通过 type 恢复不同资源
func UpdateTrashbin(c *gin.Context) {
	log.Info("project updateTrashbin funcation call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	var req EditTrashbinRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 发送请求
	_, err = service.ProjectClient.UpdateTrashbin(context.Background(), &pbp.EditTrashbinRequest{
		Id:   uint32(id),
		Type: req.Type,
	})
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
