package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/kelseyhightower/envconfig"

	"github.com/aspirin100/TaskManager/internal/database"
	"github.com/aspirin100/TaskManager/internal/logger"
	taskUsecase "github.com/aspirin100/TaskManager/internal/usecase"
)

type Config struct {
	PostgresDSN string `envconfig:"TASK_SERVER_POSTGRES_DSN" default:"postgres://postgres:postgres@localhost:5432/task-manager?sslmode=disable"` //nolint:lll
	Hostname    string `envconfig:"TASK_SERVER_HOSTNAME" default:":8000"`
	Environment string `envconfig:"TASK_SERVER_ENV" default:"local"`
}

func main() {

	config := Config{}

	err := envconfig.Process("task-server", &config)
	if err != nil {
		println("configuration reading error")
		os.Exit(1)
	}

	
	logg := logger.Default()
	logger.SetupLogger(config.Environment)

	logg.Debug("logger setuped", slog.String("env", config.Environment))

	db, err := database.UpDatabase("postgres", config.PostgresDSN)
	if err != nil {
		logg.Error(err.Error())
		os.Exit(1)
	}

	handler := taskUsecase.UsecaseHandler{
		DBRepo: database.PostgresRepo{
			DB: db,
		},
	}
	_ = handler

	err = http.ListenAndServe(config.Hostname, nil)
	if err != nil {
		logg.Error(err.Error())
		os.Exit(1)
	}

}
