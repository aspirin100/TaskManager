package tasks

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	TaskID      uuid.UUID  `json:"id,omitempty"`
	Type        string     `json:"type,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description"`
	Status      uint8      `json:"status"`
	CreatedAt   time.Time  `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

type CreateTaskRequest struct {
	UserID      uuid.UUID `json:"userid,omitempty"`
	Type        string    `json:"type,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description"`
	Status      uint8     `json:"status"`
}
