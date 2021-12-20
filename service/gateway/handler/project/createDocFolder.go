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

// CreateDocFolder ... 建立文档夹
// 更新：删除 fatherType，类型根据 father_id 判定，新增 childrenPositionIndex
// @Summary creates a doc folder api
// @Description 新建文档夹
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body CreateFolderRequest true "create_folder_request"
// @Param project_id query int true "project_id"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /folder/docfolder [post]
func CreateDocFolder(c *gin.Context) {
	log.Info("project createDocFolder function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req CreateFolderRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	userID := c.MustGet("userID").(uint32)
	projectID := c.MustGet("projectID").(uint32)

	createDocFolderReq := &pbp.CreateFolderRequest{
		Name:                  req.Name,
		CreatorId:             userID,
		ProjectId:             projectID,
		FatherId:              req.FatherId,
		ChildrenPositionIndex: req.ChildrenPositionIndex,
	}

	_, err := service.ProjectClient.CreateDocFolder(context.Background(), createDocFolderReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
