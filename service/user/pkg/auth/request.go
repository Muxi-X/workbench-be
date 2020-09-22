package auth

import (
	"errors"

	"muxi-workbench-user/util"
)

type OauthTokenResponse struct {
	util.BasicResponse
	Data TokenItem `json:"data"`
}

type TokenItem struct {
	AccessToken    string `json:"access_token"`
	AccessExpired  int64  `json:"access_expired"`
	RefreshToken   string `json:"refresh_token"`
	RefreshExpired int64  `json:"refresh_expired"`
}

type AuthUserInfoResponse struct {
	util.BasicResponse
	Data UserItem `json:"data"`
}

type UserItem struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Group     string `json:"group"`
	Timejoin  string `json:"timejoin"`
	Timeleft  string `json:"timeleft"`
	RoleID    uint64 `json:"role_id"`
	Left      bool   `json:"left"`
	Info      string `json:"info"`
	AvatarURL string `json:"avatar_url"`
}

// RegisterRequest send register request.
func RegisterRequest(name, email, password string) error {
	bodyData := map[string]string{"username": name, "email": email, "password": password}

	body, err := util.SendHTTPRequest(RegisterURL, "POST", &util.RequestData{
		BodyData:    bodyData,
		ContentType: util.JsonData,
	})
	if err != nil {
		return err
	}

	var rp util.Response
	if err := util.UnmarshalBodyForCustomData(body, &rp); err != nil {
		return err
	}

	if rp.Code != 0 {
		return errors.New(rp.Message)
	}
	return nil
}

// GetTokenRequest send access token request.
func GetTokenRequest(code, clientID, clientSecret string) (*TokenItem, error) {
	query := map[string]string{"client_id": clientID, "response_type": "token", "grant_type": "authorization_code"}
	bodyData := map[string]string{"code": code, "client_secret": clientSecret}

	body, err := util.SendHTTPRequest(OauthTokenURL, "POST", &util.RequestData{
		Query:       query,
		BodyData:    bodyData,
		ContentType: util.FormData,
	})
	if err != nil {
		return nil, err
	}

	var rp OauthTokenResponse
	if err := util.UnmarshalBodyForCustomData(body, &rp); err != nil {
		return nil, err
	}

	if rp.Code != 0 {
		return nil, errors.New(rp.Message)
	}

	if rp.Data.AccessToken == "" {
		return nil, errors.New("Getting failed, access token is blank")
	}
	return &rp.Data, nil
}

// RefreshTokenRequest send refresh token request.
func RefreshTokenRequest(refreshToken, clientID, clientSecret string) (*TokenItem, error) {
	query := map[string]string{"client_id": clientID, "grant_type": "refresh_token"}
	bodyData := map[string]string{"refresh_token": refreshToken, "client_secret": clientSecret}

	body, err := util.SendHTTPRequest(OauthRefreshURL, "POST", &util.RequestData{
		Query:       query,
		BodyData:    bodyData,
		ContentType: util.FormData,
	})
	if err != nil {
		return nil, err
	}

	var rp OauthTokenResponse
	if err := util.UnmarshalBodyForCustomData(body, &rp); err != nil {
		return nil, err
	}

	if rp.Code != 0 {
		return nil, errors.New(rp.Message)
	}

	if rp.Data.AccessToken == "" {
		return nil, errors.New("Getting failed, access token is blank")
	}
	return &rp.Data, nil
}

// GetInfoRequest send user info request.
func GetInfoRequest(token string) (*UserItem, error) {
	body, err := util.SendHTTPRequest(UserInfoURL, "GET", &util.RequestData{
		Header: map[string]string{"token": token},
	})
	if err != nil {
		return nil, err
	}

	var rp AuthUserInfoResponse
	if err := util.UnmarshalBodyForCustomData(body, &rp); err != nil {
		return nil, err
	}

	if rp.Code != 0 {
		return nil, errors.New(rp.Message)
	}
	return &rp.Data, nil
}
