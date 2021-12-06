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

// GetFileChildren ... 获取某个文件夹下的文件树
// @Summary get file children api
// @Description 获取某个文件夹下的节点(文件或文件夹)
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param id path int true "file_folder_id"
// @Param project_id query int true "project_id"
// @Success 200 {object} GetFileChildrenResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /folder/file_children/{id} [get]
func GetFileChildren(c *gin.Context) {
	log.Info("project getFileTree function call.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 folderID
	folderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}
	// TODO 验证projectId
	// 发送请求
	getFileTreeResp, err := service.ProjectClient.GetFileChildren(context.Background(), &pbp.GetRequest{
		Id: uint32(folderID),
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 解析结果
	var list []*ChildrenInfo
	for _, child := range getFileTreeResp.List {
		list = append(list, &ChildrenInfo{
			Type:        child.Type,
			Name:        child.Name,
			CreatTime:   child.CreatTime,
			CreatorName: child.CreatorName,
			Path:        child.Path,
		})
	}

	SendResponse(c, nil, &GetFileChildrenResponse{
		Count:        uint32(len(list)),
		FileChildren: list,
	})
}
