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
	tasksRepository "github.com/aspirin100/TaskManager/internal/tasks/repository"
	"github.com/aspirin100/TaskManager/internal/tasks/service/parser"
	"github.com/aspirin100/TaskManager/internal/tasks/service/response"
)

type ctxKey struct{}

var CtxUserIDKey ctxKey

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

			userID, err := parser.ParseUserID(log, r)
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

			ctx := context.WithValue(r.Context(), CtxUserIDKey, userID.String())

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
