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

// GetDocChildren ... 获取某个文档夹下的文档树
// @Summary get doc children api
// @Description 获取某个文档夹下的节点(文档或文档夹)
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param id path int true "doc_folder_id"
// @Param project_id query int true "project_id"
// @Success 200 {object} GetFileChildrenResponse
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /folder/doc_children/{id} [get]
func GetDocChildren(c *gin.Context) {
	log.Info("project getDocTree function call.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 folderID
	folderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 发送请求
	getDocTreeResp, err := service.ProjectClient.GetDocChildren(context.Background(), &pbp.GetRequest{
		Id: uint32(folderID),
	})
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 解析结果
	list := FormatChildren(getDocTreeResp.Children)

	// 返回结果
	SendResponse(c, nil, &GetFileChildrenResponse{
		Count:        uint32(len(list)),
		FileChildren: list,
	})
}
