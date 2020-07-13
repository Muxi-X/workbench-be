package auth

import (
	"github.com/spf13/viper"
)

const (
	authBasicPath    = "/auth/api"
	registerPath     = "/signup"
	userInfoPath     = "/user"
	authPath         = "/oauth"
	tokenPath        = "/oauth/token"
	refreshTokenPath = "/oauth/token/refresh"
	clientStorePath  = "/oauth/store"

	defaultAuthHost = "localhost:8083"
)

var (
	// muxi-auth-server request url
	RegisterURL     string
	UserInfoURL     string
	OauthURL        string
	OauthTokenURL   string
	OauthRefreshURL string
	ClientStoreURL  string

	clientID     string
	clientSecret string
)

func InitVar() {
	authHost := viper.GetString("auth_server.host")
	if authHost == "" {
		authHost = defaultAuthHost
	}

	basicURL := "http://" + authHost + authBasicPath
	RegisterURL = basicURL + registerPath
	UserInfoURL = basicURL + userInfoPath
	OauthURL = basicURL + authPath
	OauthTokenURL = basicURL + tokenPath
	OauthRefreshURL = basicURL + refreshTokenPath
	ClientStoreURL = basicURL + clientStorePath

	clientID = viper.GetString("auth_server.client_id")
	clientSecret = viper.GetString("auth_server.client_secret")
}
