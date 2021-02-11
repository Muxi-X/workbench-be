package project

import (
	"context"
	"fmt"
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

// UpdateFileTree ... 修改文件树
func UpdateFileTree(c *gin.Context) {
	log.Info("Project filetree Update function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 projectID
	fileID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取请求
	var req UpdateFileTreeRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 处理请求
	var item string
	var children string
	for _, v := range req.FileTree {
		if v.Type {
			item = fmt.Sprintf("%s-%d", v.Id, 1)
		} else {
			item = fmt.Sprintf("%s-%d", v.Id, 0)
		}
		if children == "" {
			children = item
			continue
		}
		children = fmt.Sprintf("%s,%s", children, item)
	}

	// 构造请求
	_, err = service.ProjectClient.UpdateFileTree(context.Background(), &pbp.UpdateTreeRequest{
		Id:   uint32(fileID),
		Tree: children,
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
