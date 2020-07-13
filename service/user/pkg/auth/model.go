package auth

import (
	"errors"
	"time"
)

type OauthManagerModel struct {
	accessToken    string
	accessExpired  int64 // 过期时间，时间戳
	refreshToken   string
	refreshExpired int64
	clientID       string
	clientSecret   string
}

var OauthManager *OauthManagerModel

func (m *OauthManagerModel) Init() {
	OauthManager = new(OauthManagerModel)

	if clientID != "" && clientSecret != "" {
		OauthManager.SetClient(clientID, clientSecret)
	}
}

func (m *OauthManagerModel) SetClient(clientID, clientSecret string) {
	m.clientID = clientID
	m.clientSecret = clientSecret
}

func (m *OauthManagerModel) checkTokenValid() bool {
	return m.accessExpired > time.Now().Unix()
}

func (m *OauthManagerModel) checkRefreshTokenValid() bool {
	return m.refreshExpired > time.Now().Unix()
}

func (m *OauthManagerModel) GetAccessToken() (string, error) {
	if ok := m.checkTokenValid(); !ok {
		if ok := m.checkRefreshTokenValid(); !ok {
			return "", errors.New("Refresh token has expired, auth again.")
		}
		if err := m.RefreshAccessToken(); err != nil {
			return "", err
		}
	}

	return m.accessToken, nil
}

// ExchangeAccessTokenWithCode gets access token with code.
func (m *OauthManagerModel) ExchangeAccessTokenWithCode(code string) error {
	if err := m.checkClientInfo(); err != nil {
		return err
	}

	item, err := GetTokenRequest(code, m.clientID, m.clientSecret)
	if err != nil {
		return err
	}

	now := time.Now()
	m.accessToken = item.AccessToken
	m.refreshToken = item.RefreshToken
	m.accessExpired = convertToUnixTime(item.AccessExpired, now)
	m.refreshExpired = convertToUnixTime(item.RefreshExpired, now)

	return nil
}

// RefreshAccessToken refresh access token.
func (m *OauthManagerModel) RefreshAccessToken() error {
	if err := m.checkClientInfo(); err != nil {
		return err
	}

	item, err := RefreshTokenRequest(m.refreshToken, m.clientID, m.clientSecret)
	if err != nil {
		return err
	}

	now := time.Now()
	m.accessToken = item.AccessToken
	m.refreshToken = item.RefreshToken
	m.accessExpired = convertToUnixTime(item.AccessExpired, now)
	m.refreshExpired = convertToUnixTime(item.RefreshExpired, now)

	return nil
}

func (m *OauthManagerModel) checkClientInfo() error {
	if m.clientID == "" || m.clientSecret == "" {
		return errors.New("client info is blank")
	}
	return nil
}

func convertToUnixTime(t int64, now time.Time) int64 {
	expire := time.Duration(t * int64(time.Second))
	return now.Add(expire).Unix()
}
