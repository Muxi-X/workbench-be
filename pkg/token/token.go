package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var (
	jwtKey string

	ErrTokenInvalid = errors.New("The token is invalid.")
	ErrTokenExpired = errors.New("The token is expired.")
)

// getJwtKey get jwtKey.
func getJwtKey() string {
	if jwtKey == "" {
		jwtKey = viper.GetString("jwt_secret")
	}
	return jwtKey
}

// TokenClaims means a claim segment in a JWT.
type TokenClaims struct {
	jwt.StandardClaims
	ID     uint32 `json:"id"`
	Role   uint32 `json:"role"`
	TeamID uint32 `json:"team_id"`
}

// TokenPayload is a required payload when generates token.
type TokenPayload struct {
	ID      uint32        `json:"id"`
	Role    uint32        `json:"role"`
	TeamID  uint32        `json:"team_id"`
	Expired time.Duration `json:"expired"` // 有效时间
}

// TokenResolve means returned payload when resolves token.
type TokenResolve struct {
	ID     uint32 `json:"id"`
	Role   uint32 `json:"role"`
	TeamID uint32 `json:"team_id"`
	// ExpiresAt int64  `json:"expires_at"` // 过期时间（时间戳，10位）
}

// GenerateToken generates token.
func GenerateToken(payload *TokenPayload) (string, error) {
	claims := &TokenClaims{
		ID:     payload.ID,
		Role:   payload.Role,
		TeamID: payload.TeamID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(payload.Expired.Seconds()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(getJwtKey()))
}

// ResolveToken resolves token.
func ResolveToken(tokenStr string) (*TokenResolve, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(getJwtKey()), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrTokenInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrTokenInvalid
	}

	exp, ok := claims["exp"]
	if !ok {
		return nil, errors.New("The exp is not in token.")
	}

	// 校验有效时间
	expiresAt := int64(exp.(float64))
	if expiresAt <= time.Now().Unix() {
		return nil, ErrTokenExpired
	}

	id, ok := claims["id"]
	if !ok {
		return nil, errors.New("The id is not in token.")
	}

	role, ok := claims["role"]
	if !ok {
		return nil, errors.New("The role is not in token.")
	}

	teamID, ok := claims["team_id"]
	if !ok {
		return nil, errors.New("The teamID is not in token.")
	}

	t := &TokenResolve{
		ID:     uint32(id.(float64)),
		Role:   uint32(role.(float64)),
		TeamID: uint32(teamID.(float64)),
	}
	return t, nil
}
