package iplimiter

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/time/rate"
)

type Config struct {
	Rps   rate.Limit `envconfig:"RPS" required:"true"`
	Burst int        `envconfig:"BURST" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("IP_LIMITER", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		panic(err)
	}

	return config
}
