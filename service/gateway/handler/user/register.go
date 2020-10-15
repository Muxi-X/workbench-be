package user

import (
	"context"

	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pb "muxi-workbench-user/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// register ... 注册
// @Summary register api
// @Description register user
// @Tags auth
// @Accept  application/json
// @Produce  application/json
// @Param object body RegisterRequest false "register_request"
// @Security ApiKeyAuth
// @Success 200 {object} handler.Response
// @Router /auth/signup [post]
func Register(c *gin.Context) {
	log.Info("User register function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取请求
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造请求给 register
	registerReq := &pb.RegisterRequest{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	// 注册
	// TO DO: 判断用户已存在，错误
	_, err := service.UserClient.Register(context.Background(), registerReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
