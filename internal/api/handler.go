package handle

import (
	"net/http"
	"time"
)

type handler interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
	GetTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
}

type Task struct {
	ID          int32      `json:"id,omitempty"`
	Type        string     `json:"type"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      uint8      `json:"status"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type CreateTaskResponse struct {
	ID     int64   `json:"id"`
	Status uint8   `json:"status"`
	Error  *string `json:"error,omitempty"`
}
