package response

import (
	"net/http"

	"github.com/go-chi/render"
)

const (
	StatusOK    = "OK"
	StatusError = "Error"
	ErrNilString = ""
)

type Response struct {
	TaskID string  `json:"taskid,omitempty"`
	Status string  `json:"status"`
	Error  *string `json:"error,omitempty"`
}

func Error(msg string, taskID string) Response {
	return Response{
		TaskID: taskID,
		Status: StatusError,
		Error:  &msg,
	}
}

func ResponseOK(w http.ResponseWriter, r *http.Request, taskID string) {
	render.JSON(w, r, Response{
		TaskID: taskID,
		Status: StatusOK,
	})
}
