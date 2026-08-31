package dbpostgres

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config holds the PostgreSQL connection settings loaded from environment
// variables with the POSTGRES prefix. It includes connection parameters
// and the operation timeout for database queries.
type Config struct {
	Host     string        `envconfig:"HOST" required:"true"`
	Port     string        `envconfig:"PORT" default:"5432"`
	Username string        `envconfig:"USER" required:"true"`
	Password string        `envconfig:"PASSWORD" required:"true"`
	DB       string        `envconfig:"DB" required:"true"`
	Timeout  time.Duration `envconfig:"TIMEOUT" required:"true"`
}

// NewConfig loads the database configuration from environment variables.
// It expects variables with the POSTGRES prefix (e.g., POSTGRES_HOST,
// POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB,
// POSTGRES_TIMEOUT).
//
// Returns an error if required environment variables are missing or
// have invalid values.
func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("POSTGRES", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}

	return config, nil
}

// NewConfigMust loads the database configuration from environment variables
// and panics if loading fails. This is useful for application startup where
// configuration errors should prevent the application from starting.
func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get postgres connection pool config: %w", err)
		panic(err)
	}

	return config
}
