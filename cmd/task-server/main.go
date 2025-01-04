package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kelseyhightower/envconfig"

	"github.com/aspirin100/TaskManager/internal/api/server/middleware/logger"
	"github.com/aspirin100/TaskManager/internal/logger/sl"
	tasksUsecase "github.com/aspirin100/TaskManager/internal/tasks/handlers"
	tasksRepository "github.com/aspirin100/TaskManager/internal/tasks/repository"
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

	db, err := tasksRepository.New(config.PostgresDSN)
	if err != nil {
		logg.Error(err.Error())
		os.Exit(1)
	}

	router := chi.NewRouter()

	router.Use(logger.New(logg))
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)

	router.Post("/{userID}/task", tasksUsecase.CreateNewTask(logg, &db))
	router.Get("/{userID}/task", tasksUsecase.GetTask(logg, &db))

	server := http.Server{
		Addr:    config.Hostname,
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		logg.Error("server start failed", sl.Err(err))
		os.Exit(1)
	}

	logg.Error("server stopped")
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
