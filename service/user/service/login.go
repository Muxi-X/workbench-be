package service

import (
	"context"
	"fmt"

	"muxi-workbench-user/errno"
	model "muxi-workbench-user/model"
	pb "muxi-workbench-user/proto"
	"muxi-workbench-user/util"
	e "muxi-workbench/pkg/err"
)

var (
	accessToken  string
	refreshToken string

	authAddr        string
	authBasicRouter = "/auth/api/v1"
	registerPath    = "/signup"
	authPath        = "/auth"
	tokenPath       = "/token"
	refreshPath     = "/token/refresh"

	oauthURL = ""
)

// Login ... 登录
// 如果无 code，则返回 oauth 的地址，让前端去请求 oauth，
// 否则，用 code 获取 oauth 的 access token，并生成 auth token，返回给前端。
func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest, res *pb.LoginResponse) error {
	if req.OauthCode == "" {
		res.RedirectUrl = oauthURL
		return nil
	}

	// 获取 access token
	if err := exchangeAccessTokenWithCode(req.OauthCode); err != nil {
		return e.ServerErr(errno.ErrAccessToken, err.Error())
	}

	// 根据 eamil 查询 user_id
	user, err := model.GetUserByEmail(req.Email)
	if err != nil || user == nil {
		return e.ServerErr(errno.ErrBadRequest, "email error")
	}

	// 生成 auth token
	token, err := util.GenerateToken(user.ID)
	if err != nil {
		return e.ServerErr(errno.ErrAuthToken, err.Error())
	}
	res.Token = token

	return nil
}

type OauthTokenResponse struct {
	Code    int
	Message string
	Data    *TokenItem
}

type TokenItem struct {
	AccessToken    string
	AccessExpired  int64
	RefreshToken   string
	RefreshExpired int64
}

func exchangeAccessTokenWithCode(code string) error {
	url := "http://localhost:8083/auth/api/oauth/token"
	query := map[string]string{"client_id": clientID, "response_type": "token", "grant_type": "authorization_code"}
	formData := map[string]string{"code": code, "client_secret": clientSecret}

	body, err := util.SendHTTPRequest(url, "POST", &util.RequestData{
		Query:       query,
		BodyData:    formData,
		ContentType: util.FormData,
	})
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	var rp OauthTokenResponse
	if err := util.MarshalBodyForCustomData(body, &rp); err != nil {
		return err
	}
	fmt.Println(rp)

	accessToken = rp.Data.AccessToken
	refreshToken = rp.Data.RefreshToken

	return nil
}
