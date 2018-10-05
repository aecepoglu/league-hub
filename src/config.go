package main

import "github.com/caarlos0/env"

type ConfigType struct {
	RedisUri string `env:"REDIS_URI,required"`
}

var config ConfigType

func LoadConfig() error {
	err := env.Parse(&config)
	return err
}
