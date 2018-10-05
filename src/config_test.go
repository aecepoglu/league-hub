package main

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

func resetConf() {
	config.RedisUri = "old val"
}

func TestLoadConfig(t *testing.T) {
	resetConf();

	os.Setenv("REDIS_URI", "my redis uri")
	defer os.Unsetenv("REDIS_URI")

	assert.Nil(t, LoadConfig())

	assert.Equal(t, config.RedisUri, "my redis uri")
}

func TestLoadConfigFail(t *testing.T) {
	resetConf();

	assert.NotNil(t, LoadConfig())
}
