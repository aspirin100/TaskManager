package handle

import (
	"context"

	"github.com/aspirin100/TaskManager/internal/tasks"
	"github.com/google/uuid"
)

type handler interface {
	CreateTask(ctx context.Context, req CreateTaskRequest) (CreateTaskResponse, error)
	GetTask(ctx context.Context, req GetTaskRequest) (tasks.Task, error)
	UpdateTask(ctx context.Context, req UpdateTaskRequest) (UpdateTaskResponse, error)
	DeleteTask(ctx context.Context, req DeleteTaskRequest) (DeleteTaskResponse, error)
}

type CreateTaskRequest struct {
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
	Status      uint8  `json:"status"`
}

type CreateTaskResponse struct {
	ID     uuid.UUID `json:"id"`
	Status uint8     `json:"status"`
	Error  *string   `json:"error,omitempty"`
}

type GetTaskRequest struct {
	ID uuid.UUID `json:"id"`
}

type UpdateTaskRequest struct {
	ID          uuid.UUID `json:"id"`
	Type        string    `json:"type,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description"`
	Status      uint8     `json:"status"`
}

type UpdateTaskResponse struct {
	ID      uuid.UUID `json:"id"`
	Message string    `json:"message"`
	Error   *string   `json:"error,omitempty"`
}

type DeleteTaskRequest struct {
	ID uuid.UUID `json:"id"`
}

type DeleteTaskResponse struct {
	Message string  `json:"message"`
	Error   *string `json:"error,omitempty"`
}
