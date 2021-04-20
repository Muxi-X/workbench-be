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

// UpdateProjectChildren ... 修改项目树 包括文件文档
// 用于和其他 api 配合实现移动文件
// 禁用
func UpdateProjectChildren(c *gin.Context) {
	log.Info("project updateProjectTree funcation call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 projectID
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	var req UpdateProjectChildrenRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 处理请求
	var item string
	var children string
	for _, v := range req.FileChildren {
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

	// 发送请求
	_, err = service.ProjectClient.UpdateProjectChildren(context.Background(), &pbp.UpdateProjectChildrenRequest{
		Id:       uint32(projectID),
		Children: children,
		Type:     req.Type,
	})
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
