package taskUsecase

import (
	"context"

	tasks_repo "github.com/aspirin100/TaskManager/internal/repository"
	"github.com/aspirin100/TaskManager/internal/tasks"
)

// for full inteface implementation check
var _ usecaseHandler = UsecaseHandler{}

type UsecaseHandler struct {
	DBRepo tasks_repo.PostgresRepo
}

func (h UsecaseHandler) CreateTask(ctx context.Context, req CreateTaskRequest) (CreateTaskResponse, error) {

	return CreateTaskResponse{}, nil
}

func (h UsecaseHandler) GetTask(ctx context.Context, req GetTaskRequest) (tasks.Task, error) {
	return tasks.Task{}, nil
}

func (h UsecaseHandler) UpdateTask(ctx context.Context, req UpdateTaskRequest) (UpdateTaskResponse, error) {
	return UpdateTaskResponse{}, nil
}

func (h UsecaseHandler) DeleteTask(ctx context.Context, req DeleteTaskRequest) (DeleteTaskResponse, error) {
	return DeleteTaskResponse{}, nil
}
