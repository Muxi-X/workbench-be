package model

import (
	"muxi-workbench/log"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Redis struct {
	Self *redis.Client
}

type PubSub struct {
	Self *redis.PubSub
}

var (
	RedisDB      *Redis
	PubSubClient *PubSub
)

func (*Redis) Init() {
	RedisDB = &Redis{
		Self: OpenRedisClient(),
	}
}

func (*Redis) Close() {
	_ = RedisDB.Self.Close()
}

func (*PubSub) Init(channel string) {
	PubSubClient = &PubSub{
		Self: OpenRedisPubSubClient(channel),
	}
}

func (*PubSub) Close() {
	_ = PubSubClient.Self.Close()
}

func OpenRedisClient() *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       0,
	})

	if _, err := r.Ping().Result(); err != nil {
		log.Fatal("Open redis failed", zap.String("reason", err.Error()))
	}
	return r
}

func OpenRedisPubSubClient(channel string) *redis.PubSub {
	return OpenRedisClient().Subscribe(channel)
}
