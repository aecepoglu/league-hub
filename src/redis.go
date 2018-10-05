package main

import (
	"github.com/go-redis/redis"
)


var redisClient *redis.Client

func ConnectRedis() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr: config.RedisUri,
	})

	_, err := redisClient.Ping().Result()

	return err
}
