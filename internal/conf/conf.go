// Package conf Configuration service generator
package conf

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

// Configuration config from app
type Configuration struct {
	PostgresConnString string `env:"POSTGRES_CONN_STRING"`
}

// NewConfiguration Constructor
func NewConfiguration() (*Configuration, error) {
	conf := Configuration{}
	err := env.Parse(&conf)
	if err != nil {
		return nil, fmt.Errorf("configuration /NewNewConfiguration : %v", err)
	}
	return &conf, nil
}
