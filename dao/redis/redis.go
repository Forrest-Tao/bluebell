package redis

import (
	"bluebell/setting"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func Init(cfg *setting.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})
	_, err = client.Ping().Result()
	return
}

func Close() (err error) {
	if err = client.Close(); err != nil {
		zap.L().Error("client.Close(): ", zap.Error(err))
		return
	}
	return
}
