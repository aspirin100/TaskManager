package tasksRepository_test

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"

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

func TestCreateTaskFail(t *testing.T) {

	rp, err := OpenDb()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	params := tasks.CreateTaskRequest{
		Description: "test description",
		Status:      1,
	}

	_, err = rp.CreateTask(context.Background(), params)
	if err != nil {
		log.Println(err)
	}
}

func TestCreateTask(t *testing.T) {
	rp, err := OpenDb()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	params := tasks.CreateTaskRequest{
		UserID:      uuid.MustParse("e05fa11d-eec3-4fba-b223-d6516800a047"),
		Description: "test description",
		Status:      1,
	}

	_, err = rp.CreateTask(context.Background(), params)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
}

func TestDeleteTaskFail(t *testing.T) {
	rp, err := OpenDb()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	params := tasks.CommonTaskRequest{
		TaskID: uuid.MustParse("c436ce0a-7bf8-420a-8ea2-ca798689f14e"),
		UserID: uuid.New(),
	}

	err = rp.DeleteTask(context.Background(), params)
	if err != nil {
		log.Println(err)
	}
}

func TestUpdateTask(t *testing.T) {
	rp, err := OpenDb()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	params := tasks.UpdateTaskRequest{
		TaskID:      uuid.MustParse("da405c59-bdf5-4483-9ce1-0187ebfd16a7"),
		UserID:      uuid.Nil,
		Name:        "test name",
		Description: "updated description",
		Status:      3,
	}

	_, err = rp.UpdateTask(context.Background(), params)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
}

func TestGetTaskFail(t *testing.T) {
	rp, err := OpenDb()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	params := tasks.CommonTaskRequest{
		TaskID: uuid.MustParse("da405c59-bdf5-4483-9ce1-0187ebfd16a7"),
		UserID: uuid.Nil,
	}

	fetchedTask, err := rp.GetTask(context.Background(), params)
	if err != nil {
		log.Println(err)
	}

	spew.Dump(fetchedTask)
}
