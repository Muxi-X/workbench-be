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

// Login ... 登录
// @Summary login api
// @Description login the workbench
// @Tags auth
// @Accept  application/json
// @Produce  application/json
// @Param object body LoginRequest true "login_request"
// @Security ApiKeyAuth
// @Success 200 {object} LoginResponse
// @Router /auth/login [post]
func Login(c *gin.Context) {
	log.Info("User login function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取 oauth_code
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造请求给 login
	loginReq := &pb.LoginRequest{
		OauthCode: req.OauthCode,
	}

	// 发送请求
	loginResp, err := service.UserClient.Login(context.Background(), loginReq)
	if err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
		return
	}

	// 构造返回 response
	resp := LoginResponse{
		Token:       loginResp.Token,
		RedirectURL: loginResp.RedirectUrl,
	}

	SendResponse(c, nil, resp)
}
