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

// UpdateTrashbin ... 用于恢复资源
// 通过 type 恢复不同资源
// @Summary update trashbin api
// @Description 恢复回收站资源
// @Tags project
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param id path int true "回收站中文件的 id"
// @Param object body RemoveTrashbinRequest true "remove_trashbin_request"
// @Param project_id query int true "此回收站对应的项目 id"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /trashbin/{id} [put]
func UpdateTrashbin(c *gin.Context) {
	log.Info("project updateTrashbin function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	var req RemoveTrashbinRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 发送请求
	_, err = service.ProjectClient.UpdateTrashbin(context.Background(), &pbp.RemoveTrashbinRequest{
		Id:                    uint32(id),
		Type:                  req.Type,
		FatherId:              req.FatherId,
		ChildrenPositionIndex: req.ChildrenPositionIndex,
		IsFatherProject:       req.IsFatherProject,
	})
	if err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
