package tasksUsecase

import (
	"log/slog"
	"net/http"

	"github.com/aspirin100/TaskManager/internal/logger/sl"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func ParseUserID(log *slog.Logger, r *http.Request) (uuid.UUID, error) {

	userID, err := uuid.Parse(chi.URLParam(r, "userID"))
	if err != nil {
		log.Error("user id parsing error", sl.Err(err))

		return uuid.Nil, err
	}

	return userID, nil
}
