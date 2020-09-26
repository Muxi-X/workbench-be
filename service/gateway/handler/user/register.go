package user

import (
	"context"
	// "strconv"

	"go.uber.org/zap"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	// "muxi-workbench-gateway/pkg/token"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"
	pb "muxi-workbench-user/proto"

	"github.com/gin-gonic/gin"
)

// Register 注册 api
func Register(c *gin.Context) {
	log.Info("User register function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取请求
	var req registerRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造请求给 register
	registerReq := &pb.RegisterRequest{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	// 发送请求
	_, err2 := service.UserClient.Register(context.Background(), registerReq)
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err2.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}
