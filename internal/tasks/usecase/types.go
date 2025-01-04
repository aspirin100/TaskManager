package tasksUsecase

import (
	"github.com/google/uuid"
)

type CreateTaskRequest struct {
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
	Status      uint8  `json:"status"`
}


type UpdateTaskRequest struct {
	ID          uuid.UUID `json:"id"`
	Type        string    `json:"type,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description"`
	Status      uint8     `json:"status"`
}

type DeleteTaskRequest struct {
	ID uuid.UUID `json:"id"`
}
