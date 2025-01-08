package tasksRepository_test

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/aspirin100/TaskManager/internal/tasks"
	tasksRepository "github.com/aspirin100/TaskManager/internal/tasks/repository"
)

func OpenDb() (tasksRepository.PostgresRepo, error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/task-manager?sslmode=disable")
	if err != nil {
		return tasksRepository.PostgresRepo{}, err
	}

	rp := tasksRepository.PostgresRepo{
		DB: db,
	}

	return rp, nil
}

func TestCreateTask(t *testing.T) {
	rp, err := OpenDb()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	cases := []struct {
		name        string
		expectedErr error
		request     tasks.CreateTaskRequest
	}{
		{
			name:        "ok case",
			expectedErr: nil,
			request: tasks.CreateTaskRequest{
				UserID:      uuid.MustParse("e05fa11d-eec3-4fba-b223-d6516800a047"),
				Name:        "test-name",
				Description: "test-descripition",
			},
		},
		{
			name:        "fail case",
			expectedErr: tasksRepository.ErrUserNotFound,
			request: tasks.CreateTaskRequest{
				UserID:      uuid.Nil,
				Name:        "test-name",
				Description: "test-descripition",
			},
		},
	}

	for _, params := range cases {
		t.Run(params.name, func(t *testing.T) {
			_, err = rp.CreateTask(context.Background(), params.request)

			require.EqualValues(t, params.expectedErr, err)
		})
	}

}

func TestDeleteTask(t *testing.T) {
	rp, err := OpenDb()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	// creating test task
	taskID, err := rp.CreateTask(context.Background(), tasks.CreateTaskRequest{
		UserID: uuid.MustParse("e05fa11d-eec3-4fba-b223-d6516800a047"),
	})
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	cases := []struct {
		name        string
		expectedErr error
		request     tasks.CommonTaskRequest
	}{
		{
			name:        "ok case",
			expectedErr: nil,
			request: tasks.CommonTaskRequest{
				UserID: uuid.MustParse("e05fa11d-eec3-4fba-b223-d6516800a047"),
				TaskID: taskID,
			},
		},
		{
			name:        "task not found case",
			expectedErr: tasksRepository.ErrTaskNotFound,
			request: tasks.CommonTaskRequest{
				UserID: uuid.MustParse("e05fa11d-eec3-4fba-b223-d6516800a047"),
				TaskID: uuid.Nil,
			},
		},
	}

	for _, params := range cases {
		t.Run(params.name, func(t *testing.T) {
			err = rp.DeleteTask(context.Background(), params.request)

			require.EqualValues(t, params.expectedErr, err)
		})
	}
}

func TestUpdateTaskFail(t *testing.T) {
	rp, err := OpenDb()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	params := tasks.UpdateTaskRequest{
		TaskID:      uuid.Nil,
		UserID:      uuid.Nil,
		Name:        "test name",
		Description: "updated description",
		Status:      3,
	}

	_, err = rp.UpdateTask(context.Background(), params)
	if err != nil {
		log.Println(err)
	}
}

func TestGetTaskFail(t *testing.T) {
	rp, err := OpenDb()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	params := tasks.CommonTaskRequest{
		TaskID: uuid.Nil,
		UserID: uuid.Nil,
	}

	fetchedTask, err := rp.GetTask(context.Background(), params)
	if err != nil {
		log.Println(err)
	}

	spew.Dump(fetchedTask)
}
