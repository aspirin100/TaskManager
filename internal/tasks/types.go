package tasks

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID  `json:"id,omitempty"`
	Type        string     `json:"type"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      uint8      `json:"status"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
