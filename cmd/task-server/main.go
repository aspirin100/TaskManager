package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kelseyhightower/envconfig"

	"github.com/aspirin100/TaskManager/internal/api/server/middleware/logger"
	tasks_repo "github.com/aspirin100/TaskManager/internal/repository"
	taskUsecase "github.com/aspirin100/TaskManager/internal/usecase"
)

type Config struct {
	PostgresDSN string `envconfig:"TASK_SERVER_POSTGRES_DSN" default:"postgres://postgres:postgres@localhost:5432/task-manager?sslmode=disable"` //nolint:lll
	Hostname    string `envconfig:"TASK_SERVER_HOSTNAME" default:":8000"`
	Environment string `envconfig:"TASK_SERVER_ENV" default:"local"`
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	config := Config{}

	err := envconfig.Process("task-server", &config)
	if err != nil {
		println("configuration reading error")
		os.Exit(1)
	}

	logg := setupLogger(config.Environment)
	logg.Debug("logger setuped", slog.String("env", config.Environment))

	db, err := tasks_repo.UpDatabase("postgres", config.PostgresDSN)
	if err != nil {
		logg.Error(err.Error())
		os.Exit(1)
	}

	handler := taskUsecase.UsecaseHandler{
		DBRepo: tasks_repo.PostgresRepo{
			DB: db,
		},
	}
	_ = handler

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(logger.New(logg))

	err = http.ListenAndServe(config.Hostname, nil)
	if err != nil {
		logg.Error(err.Error())
		os.Exit(1)
	}

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}