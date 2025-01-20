package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/aspirin100/TaskManager/internal/api/server/middleware/logger"
	validate "github.com/aspirin100/TaskManager/internal/api/server/middleware/user_validator"
	"github.com/aspirin100/TaskManager/internal/config"
	"github.com/aspirin100/TaskManager/internal/logger/sl"
	tasksService "github.com/aspirin100/TaskManager/internal/tasks/handlers"
	tasksRepository "github.com/aspirin100/TaskManager/internal/tasks/repository"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	logg := setupLogger(cfg.Environment)
	logg.Debug("logger setuped", slog.String("env", cfg.Environment))

	db, err := tasksRepository.New(cfg.PostgresDSN)
	if err != nil {
		logg.Error(err.Error())
		os.Exit(1)
	}

	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)

	router.Route("/{userID}", func(r chi.Router) {
		r.Use(validate.ValidateUser(logg, db))
		router.Use(logger.New(logg))

		r.Post("/task", tasksService.CreateNewTask(logg, db))
		r.Get("/task", tasksService.GetTask(logg, db))
		r.Put("/task", tasksService.UpdateTask(logg, db))
		r.Delete("/task", tasksService.DeleteTask(logg, db))
	})

	server := http.Server{
		Addr:         cfg.Hostname,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
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
