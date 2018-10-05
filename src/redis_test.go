package main

import (
	"testing"
	//"github.com/go-redis/redis"
	"github.com/alicebob/miniredis"
	"github.com/stretchr/testify/assert"
)

func resetRedis() {
	redisClient = nil
	config.RedisUri = "localhost:6379"
}

func TestConnectRedis(t *testing.T) {
	resetRedis()
	mini, err := miniredis.Run()
	assert.Nil(t, err)
	defer mini.Close()

	assert.Nil(t, ConnectRedis())
	assert.NotNil(t, redisClient)
}

func TestConfigRedisFail(t *testing.T) {
	resetConf()

	assert.NotNil(t, LoadConfig())
}
