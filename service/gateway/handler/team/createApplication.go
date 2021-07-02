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

// CreateApplication ... 创建申请
// @Summary create application api
// @Description 创建 application
// @Tags application
// @Accept  application/json
// @Produce  application/json
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /team/application [post]
func CreateApplication(c *gin.Context) {
	log.Info("Application create function call.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 userid
	userID := c.MustGet("userID").(uint32)

	// 构造 CreateApplication 请求
	CreateApplicationReq := &tpb.ApplicationRequest{
		UserId: userID,
	}

	// 向 CreateApplication 服务发送请求
	_, err := service.TeamClient.CreateApplication(context.Background(), CreateApplicationReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
