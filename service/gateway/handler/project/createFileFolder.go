package project

import (
	"context"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pbp "muxi-workbench-project/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateFileFolder ... 建立文件夹
// @Summary create a doc folder api
// @Description 新建文件夹
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body CreateFolderRequest true "create_folder_request"
// @Param project_id query int true "project_id"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /folder/filefolder [post]
func CreateFileFolder(c *gin.Context) {
	log.Info("project createFileFolder function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req CreateFolderRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	userID := c.MustGet("userID").(uint32)
	projectID := c.MustGet("projectID").(uint32)

	createFileFolderReq := &pbp.CreateFolderRequest{
		Name:                  req.Name,
		CreatorId:             userID,
		ProjectId:             projectID,
		FatherId:              req.FatherId,
		ChildrenPositionIndex: req.ChildrenPositionIndex,
	}

	_, err := service.ProjectClient.CreateFileFolder(context.Background(), createFileFolderReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// feed

	SendResponse(c, errno.OK, nil)
}
