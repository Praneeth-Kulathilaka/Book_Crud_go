package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var redisClient *redis.Client

func InitRedis() *redis.Client {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	_,err := redisClient.Ping(Ctx).Result()
	if err != nil {
		log.Println("Failed to connect redis: ",err)
		return nil
	}
	log.Println("Connected to Redis Successfully")
	return redisClient
}

func GetRedisClient() *redis.Client {
	return redisClient
}
