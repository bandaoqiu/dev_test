package redisx

import (
	"dev_test/pkg/configx"
	"github.com/go-redis/redis"
)

func InitRedis()*redis.Client{

	client := redis.NewClient(&redis.Options{
		Network:  "",
		Addr:     configx.Cfg.Redis.Addr,
		Password: configx.Cfg.Redis.PassWord,
		DB:       0,

	})
	return client
}
