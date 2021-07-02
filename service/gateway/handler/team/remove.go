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

// Remove ... 移除成员
// @Summary remove api
// @Description 移除成员
// @Tags team
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body RemoveRequest true "remove_request"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /team/member [delete]
func Remove(c *gin.Context) {
	log.Info("Remove team function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 teamID
	teamID := c.MustGet("teamID").(uint32)

	// 获取请求
	var req RemoveRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	removeReq := &tpb.RemoveRequest{
		UserList: req.UserList,
		TeamId:   teamID,
	}

	// 向 Remove 服务发送请求
	_, err := service.TeamClient.Remove(context.Background(), removeReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
