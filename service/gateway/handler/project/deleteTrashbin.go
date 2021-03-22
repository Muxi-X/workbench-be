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

// DeleteTrashbin ... 删除垃圾资源 传不同的 type 删不一样的资源
// type： 0-project 1-doc 2-file 3-doc folder 4-file folder
// 其中 project 和两个 folder 是递归删除
func DeleteTrashbin(c *gin.Context) {
	log.Info("project deleteTrashbin funcation call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	var req DeleteTrashbinRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 发送请求
	_, err = service.ProjectClient.DeleteTrashbin(context.Background(), &pbp.DeleteTrashbinRequest{
		Id:   uint32(id),
		Type: req.Type,
	})
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
