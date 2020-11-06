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

// Join ... 加入团队
// @Summary join team api
// @Description 加入 team
// @Tags team
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body JoinRequest true "join_request"
// @Security ApiKeyAuth
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /team/member [post]
func Join(c *gin.Context) {
	log.Info("Join team function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取请求
	var req JoinRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 获取 teamID
	teamID := c.MustGet("teamID").(uint32)

	// 构造 join 请求
	joinReq := &tpb.JoinRequest{
		UserList: req.UserList,
		TeamId:   teamID,
	}

	// 向 Join 服务发送请求
	_, err := service.TeamClient.Join(context.Background(), joinReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
