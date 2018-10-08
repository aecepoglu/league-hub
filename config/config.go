package config

import "github.com/caarlos0/env"

type ConfigType struct {
	RedisUri string `env:"REDIS_URI,required"`
	AdminPass string `env:"ADMIN_PASS" envDefault:"123456"`
}

func LoadConfig() (ConfigType, error) {
	var c ConfigType
	err := env.Parse(&c)
	return c, err
}
