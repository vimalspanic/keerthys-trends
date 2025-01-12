package utils

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/vimalspanic/keerthys-trends/config"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.GetConfig("redis_host") + ":" + config.GetConfig("redis_port"),
		Password: config.GetConfig("redis_password"),
		DB:       0,
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
}
