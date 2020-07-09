package service

import (
	"context"
	"time"

	"muxi-workbench-user/errno"
	"muxi-workbench-user/model"
	"muxi-workbench-user/pkg/auth"
	pb "muxi-workbench-user/proto"
	e "muxi-workbench/pkg/err"
	"muxi-workbench/pkg/token"
)

// Login ... 登录
// 如果无 code，则返回 oauth 的地址，让前端去请求 oauth，
// 否则，用 code 获取 oauth 的 access token，并生成该应用的 auth token，返回给前端。
func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest, res *pb.LoginResponse) error {
	if req.OauthCode == "" {
		res.RedirectUrl = auth.OauthURL
		return nil
	}

	// 根据 eamil 查询 user
	user, err := model.GetUserByEmail(req.Email)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	} else if user == nil {
		return e.ServerErr(errno.ErrBadRequest, "email error or user does not exist")
	}

	// 获取 access token
	if err := auth.OauthManager.ExchangeAccessTokenWithCode(req.OauthCode); err != nil {
		return e.ServerErr(errno.ErrAccessToken, err.Error())
	}

	// 生成 auth token
	token, err := token.GenerateToken(token.TokenPayload{
		ID:      user.ID,
		Expired: time.Hour * 24 * 30,
	})
	if err != nil {
		return e.ServerErr(errno.ErrAuthToken, err.Error())
	}

	res.Token = token
	return nil
}
