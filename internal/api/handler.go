package handle

import (
	"context"

	"github.com/aspirin100/TaskMaster/internal/postgres"
	"github.com/aspirin100/TaskMaster/internal/tasks"
)

var _ handler = Handler{}

type Handler struct {
	DBRepo postgres.PostgresRepo
}

func (h Handler) CreateTask(ctx context.Context, req CreateTaskRequest) (CreateTaskResponse, error) {

	return CreateTaskResponse{}, nil
}

func (h Handler) GetTask(ctx context.Context, req GetTaskRequest) (tasks.Task, error) {
	return tasks.Task{}, nil
}

func (h Handler) UpdateTask(ctx context.Context, req UpdateTaskRequest) (UpdateTaskResponse, error) {
	return UpdateTaskResponse{}, nil
}

func (h Handler) DeleteTask(ctx context.Context, req DeleteTaskRequest) (DeleteTaskResponse, error) {
	return DeleteTaskResponse{}, nil
}
