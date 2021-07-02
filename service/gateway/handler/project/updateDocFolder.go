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

// UpdateDocFolder ... 修改文档夹 改名
// @Summary update doc folder's name api
// @Description 修改文档夹名字
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param id path int true "文档夹 id"
// @Param object body UpdateFolderRequest true "update_folder_request"
// @Param project_id query int true "此文档夹所属项目 id"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /folder/docfolder/{id} [put]
func UpdateDocFolder(c *gin.Context) {
	log.Info("Project updateDocFolder function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 folderID 和请求
	var err error
	var folderID int

	folderID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
	}

	// 获取请求体
	var req UpdateFolderRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 updateReq 并发送请求
	updateReq := &pbp.UpdateFolderRequest{
		FolderId: uint32(folderID),
		Name:     req.Name,
	}

	_, err = service.ProjectClient.UpdateDocFolder(context.Background(), updateReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
