package model

import (
	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

var PubRdb *redis.Client
var SubRdb *redis.PubSub
var RdbChan = "Rdb"

func OpenRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB: 0,
	})
}
