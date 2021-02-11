package project

import (
	"context"
	"strconv"
	"strings"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetProjectTree ... 获取项目子树 包括文件树和文档树
func GetProjectTree(c *gin.Context) {
	log.Info("project getProjectTree function call.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 projectID
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 发送请求
	getFileTreeResp, err := service.ProjectClient.GetProjectTree(context.Background(), &pbp.GetRequest{
		Id: uint32(projectID),
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 解析结果
	var docList []*FileTreeItem
	var fileList []*FileTreeItem
	docRaw := strings.Split(getFileTreeResp.DocTree, ",")
	fileRaw := strings.Split(getFileTreeResp.FileTree, ",")
	for _, v := range docRaw {
		r := strings.Split(v, "-")
		if r[1] == "0" {
			docList = append(docList, &FileTreeItem{
				Id:   r[0],
				Type: false,
			})
		} else {
			docList = append(docList, &FileTreeItem{
				Id:   r[0],
				Type: true,
			})
		}
	}

	for _, v := range fileRaw {
		r := strings.Split(v, "-")
		if r[1] == "0" {
			fileList = append(fileList, &FileTreeItem{
				Id:   r[0],
				Type: false,
			})
		} else {
			fileList = append(fileList, &FileTreeItem{
				Id:   r[0],
				Type: true,
			})
		}
	}

	SendResponse(c, nil, &GetProjectTreeResponse{
		DocTree:  docList,
		FileTree: fileList,
	})
}
