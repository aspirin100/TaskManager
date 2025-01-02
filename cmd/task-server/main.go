package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/kelseyhightower/envconfig"

	handle "github.com/aspirin100/TaskMaster/internal/api"
	"github.com/aspirin100/TaskMaster/internal/postgres"
)

type Config struct {
	PostgresDSN string `envconfig:"TASK_SERVER_POSTGRES_DSN" default:"postgres://postgres:postgres@localhost:5432/task-manager?sslmode=disable"` //nolint:lll
	Hostname    string `envconfig:"TASK_SERVER_HOSTNAME" default:":8000"`
	Environment string `envconfig:"TASK_SERVER_ENV" default:"local"`
}

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

func main() {

	config := Config{}

	err := envconfig.Process("task-server", &config)
	if err != nil {
		println("configuration reading error")
		os.Exit(1)
	}

	logger := setupLogger(config.Environment)
	logger.Debug("logger setuped", slog.String("env", config.Environment))

	db, err := postgres.UpDatabase("postgres", config.PostgresDSN)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	handler := handle.Handler{
		DBRepo: postgres.PostgresRepo{
			DB: db,
		},
	}
	_ = handler

	err = http.ListenAndServe(config.Hostname, nil)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return logger
}
