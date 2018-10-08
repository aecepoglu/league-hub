package config

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("REDIS_URI", "my redis uri")
	os.Setenv("ADMIN_PASS", "my admin pass")
	defer os.Unsetenv("REDIS_URI")

	c, err := LoadConfig()
	assert.Nil(t, err)

	assert.Equal(t, c.RedisUri, "my redis uri")
	assert.Equal(t, c.AdminPass, "my admin pass")
}

func TestLoadConfigFail(t *testing.T) {
	_, err := LoadConfig()
	assert.NotNil(t, err)
}
