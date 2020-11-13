package util

import (
	"time"

	"github.com/spf13/viper"
)

// GetExpiredTime get token expired time from env or config file.
func GetExpiredTime() time.Duration {
	day := viper.GetInt("token.expired")
	return time.Hour * 24 * time.Duration(day)
}
