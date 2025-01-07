package validate

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	"github.com/aspirin100/TaskManager/internal/logger/sl"
	tasksUsecase "github.com/aspirin100/TaskManager/internal/tasks/handlers"
	"github.com/aspirin100/TaskManager/internal/tasks/handlers/response"
	tasksRepository "github.com/aspirin100/TaskManager/internal/tasks/repository"
)

type UserChecker interface {
	CheckUserExists(ctx context.Context, userID uuid.UUID) error
}

func ValidateUser(log *slog.Logger, userChecker UserChecker) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log := log.With(
			slog.String("component", "tasks/handlers/middleware/user_validator"),
		)

		log.Debug("user validator middleware started")

		fn := func(w http.ResponseWriter, r *http.Request) {
			spew.Dump(chi.URLParam(r, "userID"))

			userID, err := tasksUsecase.ParseUserID(log, r)
			if err != nil {
				render.JSON(w, r, response.Error("wrong user id format", response.ErrNilString))
				return
			}

			err = userChecker.CheckUserExists(r.Context(), userID)
			if err != nil {
				switch {
				case errors.Is(err, tasksRepository.ErrUserNotFound):
					render.JSON(w, r, response.Error("user not found", response.ErrNilString))
				default:
					render.JSON(w, r, response.Error("user validation error", response.ErrNilString))
				}

				log.Error("error from user validaton", sl.Err(err))

				return
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
