package model

import (
	"time"

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

// OpenRedisClient opens a redis client with the addr and password getting from env or config file
func OpenRedisClient() *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       0,
	})

	if _, err := r.Ping().Result(); err != nil {
		log.Error("Open redis failed", zap.String("reason", err.Error()))
	}
	return r
}

// OpenRedisPubSubClient opens a redis pubSub client
func OpenRedisPubSubClient(channel string) *redis.PubSub {
	return OpenRedisClient().Subscribe(channel)
}

/* ----------------------------------------------------------- */

// GetStringFromRedis gets a value by a string key
func GetStringFromRedis(key string) (string, bool, error) {
	val, err := RedisDB.Self.Get(key).Result()
	if err == redis.Nil {
		return "", false, nil
	} else if err != nil {
		return "", false, err
	}

	return val, true, nil
}

// SetStringInRedis sets a value by a key
func SetStringInRedis(key string, value interface{}, expiration time.Duration) error {
	return RedisDB.Self.Set(key, value, expiration).Err()
}

// HasExistedInRedis checks whether the key exists,
// return 0 if not exists
func HasExistedInRedis(key string) (bool, error) {
	val, err := RedisDB.Self.Exists(key).Result()
	if err != nil {
		return false, err
	}

	return val != 0, nil
}
