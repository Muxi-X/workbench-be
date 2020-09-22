package user

import (
	"context"
	//"strconv"

	"go.uber.org/zap"
	. "muxi-workbench-gateway/handler"
	"muxi-workbench-gateway/log"
	"muxi-workbench-gateway/pkg/errno"
	pb "muxi-workbench-user/proto"
	//"muxi-workbench-gateway/pkg/token"
	"muxi-workbench-gateway/service"
	"muxi-workbench-gateway/util"

	"github.com/gin-gonic/gin"
)

// Login 登录 api
func Login(c *gin.Context) {
	log.Info("User login function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取 oauth_code
	var req loginRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 构造请求给 login
	loginReq := &pb.LoginRequest{
		OauthCode: req.OauthCode,
	}

	// 发送请求
	loginResp, err2 := service.UserClient.Login(context.Background(), loginReq)
	if err2 != nil {
		SendError(c, errno.InternalServerError, nil, err2.Error())
		return
	}

	// 构造返回 response
	resp := loginResponse{
		Token:       loginResp.Token,
		RedirectURL: loginResp.RedirectUrl,
	}

	SendResponse(c, nil, resp)
}
