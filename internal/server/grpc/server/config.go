package grpcserver

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Address          string        `envconfig:"ADDR" required:"true"`
	ShutdownTimeout  time.Duration `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
	EnableReflection bool          `envconfig:"ENABLE_REFLECTION" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("GRPC", &config); err != nil {
		return Config{}, fmt.Errorf("error processing config: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get GRPC server config: %w", err)
		panic(err)
	}

	return config
}
