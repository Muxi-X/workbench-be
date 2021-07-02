package team

import (
	"context"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	tpb "muxi-workbench-team/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// DeleteApplication ... 删除申请
// @Summary delete applications api
// @Description 批量删除 applications
// @Tags application
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body DeleteApplicationRequest true "delete_application_request"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /team/application [delete]
func DeleteApplication(c *gin.Context) {
	log.Info("Application delete function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 构造请求
	var req DeleteApplicationRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造 DeleteApplication 请求
	DeleteApplicationReq := &tpb.DeleteApplicationRequest{
		ApplyList: req.ApplicationList,
	}

	// 向 DeleteApplication 服务发送请求
	_, err := service.TeamClient.DeleteApplication(context.Background(), DeleteApplicationReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
