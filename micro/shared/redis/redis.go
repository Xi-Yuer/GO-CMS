package redisDB

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
)

var (
	redisClient *redis.Client
	once        sync.Once
)

// InitRedis 初始化 Redis 客户端
func InitRedis(redisAddr string) *redis.Client {
	once.Do(func() {
		// 创建 Redis 客户端
		redisClient = redis.NewClient(&redis.Options{
			Addr: redisAddr,
			// 设置连接池大小
			PoolSize: 10,
		})

		// 确保连接到 Redis
		_, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("无法连接到 Redis: %v", err)
		}
	})

	return redisClient
}
