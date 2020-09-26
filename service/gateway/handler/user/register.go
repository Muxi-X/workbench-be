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

// Register 注册 api
func Register(c *gin.Context) {
	log.Info("User register function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取请求
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
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
		SendError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}
