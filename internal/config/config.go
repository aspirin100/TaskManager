package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgresDSN string        `envconfig:"TASK_SERVER_POSTGRES_DSN" default:"postgres://postgres:postgres@localhost:5432/task-manager?sslmode=disable"` //nolint:lll
	Hostname    string        `envconfig:"TASK_SERVER_HOSTNAME" default:":8000"`
	Timeout     time.Duration `envconfig:"TASK_SERVER_TIMEOUT" default:"5s"`
	IdleTimeout time.Duration `envconfig:"TASK_SERVER_IDLE_TIMEOUT" default:"60s"`
	Environment string        `envconfig:"TASK_SERVER_ENV" default:"local"`
}

func New() (*Config, error) {

	cfg := &Config{}

	err := envconfig.Process("task-server", cfg)
	if err != nil {
		return nil, fmt.Errorf("error during config reading: %w", err)
	}

	return cfg, nil
}
