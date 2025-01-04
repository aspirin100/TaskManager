package tasksUsecase

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	"github.com/aspirin100/TaskManager/internal/logger/sl"
	"github.com/aspirin100/TaskManager/internal/tasks"
	"github.com/aspirin100/TaskManager/internal/tasks/handlers/response"
	tasksRepository "github.com/aspirin100/TaskManager/internal/tasks/repository"
)

type TaskCreator interface {
	CreateTask(ctx context.Context, params tasks.CreateTaskRequest) (uuid.UUID, error)
}

func CreateNewTask(log *slog.Logger, taskCreator TaskCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "tasksUsecase.CreateNewTask"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req tasks.CreateTaskRequest

		userID, err := uuid.Parse(chi.URLParam(r, "userID"))
		if err != nil {
			log.Error("user id parsing error", sl.Err(err))
			render.JSON(w, r, response.Error("wrong user id format"))

			return
		}

		req.UserID = userID

		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Error("request body is empty")
				render.JSON(w, r, response.Error("empty request"))
			}

			log.Error("failed to decode request body", sl.Err(err))
			render.JSON(w, r, response.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		taskID, err := taskCreator.CreateTask(r.Context(), req)
		if err != nil {
			switch {
			case errors.Is(err, tasksRepository.ErrTaskNotFound):
				log.Error("task not found", sl.Err(err))
				render.JSON(w, r, response.Error("task not found"))
			case errors.Is(err, tasksRepository.ErrUserNotFound):
				log.Error("user not found", sl.Err(err))
				render.JSON(w, r, response.Error("user not found"))
			default:
				log.Error("create task failed", sl.Err(err))
				render.JSON(w, r, response.Error("create task failed"))
			}

			return
		}

		log.Info("task created:", slog.String("taskID", taskID.String()))

		response.ResponseOK(w, r, taskID)
	}
}


