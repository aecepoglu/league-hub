package main

import (
	"testing"
	//"github.com/go-redis/redis"
	"github.com/alicebob/miniredis"
	"github.com/stretchr/testify/assert"
)

func TestConnectRedis(t *testing.T) {
	mini, err := miniredis.Run()
	assert.Nil(t, err)
	defer mini.Close()

	r, err := connectRedis(mini.Addr())
	assert.NotNil(t, r)
	assert.Nil(t, err)
}
