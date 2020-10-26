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

// GetFileTree ... 获取文件树
func GetFileTree(c *gin.Context) {
	log.Info("project getFileTree function call.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 projectID
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 发送请求
	getFileTreeResp, err := service.ProjectClient.GetFileTree(context.Background(), &pbp.GetRequest{
		Id: uint32(projectID),
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, nil, &GetFileTreeResponse{
		FileTree: getFileTreeResp.Tree,
	})
}
