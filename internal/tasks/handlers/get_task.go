package tasksUsecase

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	"github.com/aspirin100/TaskManager/internal/logger/sl"
	"github.com/aspirin100/TaskManager/internal/tasks"
	"github.com/aspirin100/TaskManager/internal/tasks/handlers/response"
	tasksRepository "github.com/aspirin100/TaskManager/internal/tasks/repository"
)

type getTaskRequest struct {
	TaskID uuid.UUID `json:"taskid"`
}

type TaskReader interface {
	GetTask(ctx context.Context, taskID uuid.UUID) (tasks.Task, error)
}

func GetTask(log *slog.Logger, taskReader TaskReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "tasksUsecase.GetTask"

		log := log.With(
			slog.String("op", op),
			slog.String("requestID", middleware.GetReqID(r.Context())),
		)

		_, err := ParseUserID(log, r)
		if err != nil {
			render.JSON(w, r, response.Error("wrong user id format", response.ErrNilString))

			return
		}

		var req getTaskRequest

		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Error("request body is empty")
				render.JSON(w, r, response.Error("empty request", response.ErrNilString))
			} else {
				log.Error("failed to decode request body", sl.Err(err))
				render.JSON(w, r, response.Error("failed to decode request", response.ErrNilString))
			}

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		task, err := taskReader.GetTask(r.Context(), req.TaskID)
		if err != nil {
			switch {
			case errors.Is(err, tasksRepository.ErrTaskNotFound):
				log.Error("task not found", sl.Err(err))
				render.JSON(w, r, response.Error("task not found", req.TaskID.String()))
			default:
				log.Error("get task failed", sl.Err(err))
				render.JSON(w, r, response.Error("get task failed", req.TaskID.String()))
			}

			return
		}

		render.JSON(w, r, task)
	}
}
