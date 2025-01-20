package tasksService

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	validate "github.com/aspirin100/TaskManager/internal/api/server/middleware/user_validator"
	"github.com/aspirin100/TaskManager/internal/logger/sl"
	"github.com/aspirin100/TaskManager/internal/tasks"
	"github.com/aspirin100/TaskManager/internal/tasks/service/response"
	tasksRepository "github.com/aspirin100/TaskManager/internal/tasks/repository"
)

type TaskDeleter interface {
	DeleteTask(ctx context.Context, params tasks.CommonTaskRequest) error
}

func DeleteTask(log *slog.Logger, taskDeleter TaskDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "tasksUsecase.DeleteTask"

		log := log.With(
			slog.String("op", op),
			slog.String("requestID", middleware.GetReqID(r.Context())),
		)

		userID := uuid.MustParse(r.Context().Value(validate.CtxUserIDKey).(string))

		var req tasks.CommonTaskRequest

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			switch {
			case errors.Is(err, io.EOF):
				log.Error("request body is empty")
				render.JSON(w, r, response.Error("empty request", response.ErrNilString))
			default:
				log.Error("failed to decode request body", sl.Err(err))
				render.JSON(w, r, response.Error("failed to decode request", response.ErrNilString))
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		req.UserID = userID

		err = taskDeleter.DeleteTask(r.Context(), req)
		if err != nil {
			switch {
			case errors.Is(err, tasksRepository.ErrTaskNotFound):
				log.Error("task not found", sl.Err(err))
				w.WriteHeader(http.StatusNotFound)
				render.JSON(w, r, response.Error("task not found", req.TaskID.String()))
			default:
				log.Error("delete task failed", sl.Err(err))
				w.WriteHeader(http.StatusInternalServerError)
				render.JSON(w, r, response.Error("delete task failed", req.TaskID.String()))
			}

			return
		}

		response.ResponseOK(w, r, req.TaskID.String())
	}
}
