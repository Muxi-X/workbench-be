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

// DeleteDocFolder ... 删除文档夹
// @Summary delete a doc folder api
// @Description 删除文档文件夹
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param id path int true "doc_folder_id"
// @Param project_id query int true "project_id"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /folder/docfolder/{id} [delete]
func DeleteDocFolder(c *gin.Context) {
	log.Info("project deleteDocFolder function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	folderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 获取 userID
	userID := c.MustGet("userID").(uint32)
	role := c.MustGet("role").(uint32)

	projectID := c.MustGet("projectID").(uint32)
	deleteDocFolderReq := &pbp.DeleteRequest{
		Id:        uint32(folderId),
		UserId:    userID,
		Role:      role,
		ProjectId: projectID,
	}

	_, err = service.ProjectClient.DeleteDocFolder(context.Background(), deleteDocFolderReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
