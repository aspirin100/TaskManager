package response

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

type Response struct {
	TaskID uuid.UUID `json:"taskid,omitempty"`
	Status string    `json:"status"`
	Error  *string   `json:"error,omitempty"`
}

func Error(msg string, taskID uuid.UUID) Response {
	return Response{
		TaskID: taskID,
		Status: StatusError,
		Error:  &msg,
	}
}

func ResponseOK(w http.ResponseWriter, r *http.Request, taskID uuid.UUID) {
	render.JSON(w, r, Response{
		TaskID: taskID,
		Status: StatusOK,
	})
}
