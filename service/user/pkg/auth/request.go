package auth

import (
	"errors"
	"fmt"

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
	// Birthday     string `json:"birthday"`
	// Hometown     string `json:"hometown"`
	// PersonalBlog string `json:"personal_blog"`
	// Github    string `json:"github"`
	// Flickr       string `json:"flickr" column:"flickr"`
	// Weibo        string `json:"weibo" column:"weibo"`
	// Zhihu        string `json:"zhihu" column:"zhihu"`
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
	fmt.Println(string(body))

	var rp util.Response
	if err := util.MarshalBodyForCustomData(body, &rp); err != nil {
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
	if err := util.MarshalBodyForCustomData(body, &rp); err != nil {
		return nil, err
	}
	fmt.Println(rp.Data)

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
	if err := util.MarshalBodyForCustomData(body, &rp); err != nil {
		return nil, err
	}
	fmt.Println(rp.Data)

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
	fmt.Println(string(body))

	var rp AuthUserInfoResponse
	if err := util.MarshalBodyForCustomData(body, &rp); err != nil {
		return nil, err
	}

	if rp.Code != 0 {
		return nil, errors.New(rp.Message)
	}
	return &rp.Data, nil
}
