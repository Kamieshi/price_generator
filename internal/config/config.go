package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	RedisHost             string `env:"REDIS_HOST"`
	RedisPort             string `env:"REDIS_PORT"`
	CountCompany          int    `env:"COUNT_COMPANY" envDefault:"1"`
	CountUpdatePerSecond  int    `env:"COUNT_UPDATE_PER_SECOND" envDefault:"5"`
	OnlyPositiveIncrement bool   `env:"ONLY_POSITIVE_INCREMENT" envDefault:"false"`
}

func GetConfig() (*Config, error) {
	config := Config{}
	err := env.Parse(&config)
	if err != nil {
		return nil, fmt.Errorf("config / GetConfig / err parse : %v", err)
	}
	return &config, nil
}
