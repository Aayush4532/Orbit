package db

import (
	"Orbit/configs"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

func SetUpRedis() {
	cfg := configs.LoadConfig().Redis
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Url,
		Password: cfg.Password,
		DB:       0,
	})

	redisClient = client
}

func GetRedisClient() *redis.Client {
	redisOnce.Do(func() {
		SetUpRedis()
	})
	return redisClient
}
