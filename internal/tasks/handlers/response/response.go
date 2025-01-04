package response

import "github.com/google/uuid"

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

type Response struct {
	TaskID         uuid.UUID `json:"taskID,omitempty"`
	Status string    `json:"Status"`
	Error      *string   `json:"error,omitempty"`
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:      &msg,
	}
}
