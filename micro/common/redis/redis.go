package redisDB

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var redisClient *redis.Client

// InitRedis 初始化 Redis 客户端
func InitRedis(redisAddr string) *redis.Client {
	// 如果 Redis 客户端已经初始化，则直接返回
	if redisClient != nil {
		return redisClient
	}

	// 创建 Redis 客户端
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	// 确保连接到 Redis
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("无法连接到 Redis: %v", err)
	}

	return redisClient
}
