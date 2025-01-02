package handle_test

import (
	"context"
	"testing"

	handle "github.com/aspirin100/TaskManager/internal/api"
	"github.com/aspirin100/TaskManager/internal/postgres"
)

func TestCreateTask(t *testing.T) {

	db, err := postgres.UpDatabase("postgres", "postgres://postgres:postgres@localhost:5432/task-manager?sslmode=disable")
	if err != nil {
		t.Fail()
	}

	handler := handle.Handler{
		DBRepo: postgres.PostgresRepo{
			DB: db,
		},
	}

	handler.CreateTask(context.Background(), handle.CreateTaskRequest{Description: "test description", Status: 1})
}
