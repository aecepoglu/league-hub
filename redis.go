package main

import (
	"github.com/go-redis/redis"
)

func connectRedis(uri string) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: uri,
	})

	_, err := redisClient.Ping().Result()

	return redisClient, err
}
