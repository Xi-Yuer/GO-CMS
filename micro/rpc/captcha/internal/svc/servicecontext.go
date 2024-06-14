package svc

import (
	"github.com/redis/go-redis/v9"
	appConfig "micro/config"
	"micro/rpc/captcha/internal/config"
	redisDB "micro/shared/redis"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		Redis:  redisDB.InitRedis(appConfig.RedisHost),
	}
}
