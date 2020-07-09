package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var jwtKey string

// getJwtKey get jwtKey.
func getJwtKey() string {
	if jwtKey == "" {
		jwtKey = viper.GetString("jwt_key")
	}
	return jwtKey
}

type TokenClaims struct {
	ID uint32 `json:"id"`
	jwt.StandardClaims
}

type TokenPayload struct {
	ID      uint32        `json:"id"`
	Expired time.Duration `json:"expired"`
}

type TokenResolve struct {
	ID      uint32 `json:"id"`
	Expired int64  `json:"expired"`
}

// GenerateToken generate token.
func GenerateToken(payload TokenPayload) (string, error) {
	claims := &TokenClaims{
		ID: payload.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(payload.Expired.Seconds()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(getJwtKey()))
}

// ResolveToken resolve token，return id and error.
func ResolveToken(tokenStr string) (uint32, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(getJwtKey()), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		sub, ok := claims["id"]
		if !ok {
			return 0, errors.New("The id is not in token.")
		}
		return uint32(sub.(float64)), nil
	}
	return 0, errors.New("Unknown error.")
}
