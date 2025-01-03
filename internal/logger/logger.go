package logger

import (
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

var lg MLogger = slog.New(
	slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
)

type MLogger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Error(msg string, args ...any)
	Warn(msg string, args ...any)
}

func Default() MLogger {
	return lg
}

func SetupLogger(env string) {
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

	lg = logger
}

func Fatal(msg string, v ...any) {
	lg.Error(msg, v...)
	os.Exit(1)
}
