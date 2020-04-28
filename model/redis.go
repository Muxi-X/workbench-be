package model

import (
	"muxi-workbench/log"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type RedisDB struct {
	Self *redis.Client
	//PubSub *redis.PubSub
}

type PubSub struct {
	Self *redis.PubSub
}

var (
	Rdb   *RedisDB
	PSCli *PubSub
)

func (*RedisDB) Init() {
	Rdb = &RedisDB{
		Self: OpenRedisClient(),
	}
}

func (*RedisDB) Close() {
	_ = Rdb.Self.Close()
}

func (*PubSub) Init(channel string) {
	PSCli = &PubSub{
		Self: OpenRedisPubSubClient(channel),
	}
}

func (*PubSub) Close() {
	_ = PSCli.Self.Close()
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
