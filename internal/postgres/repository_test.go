package postgres_test

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/aspirin100/TaskMaster/internal/postgres"
	"github.com/google/uuid"
)

func TestInsertNewTaskFail(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/task-manager?sslmode=disable")
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	rp := postgres.PostgresRepo{
		DB: db,
	}

	params := postgres.InsertTaskParams{
		Description: "test description",
		Status:      1,
	}

	_, err = rp.InsertNewTask(context.Background(), params)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
}

func TestInsertNewTask(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/task-manager?sslmode=disable")
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	rp := postgres.PostgresRepo{
		DB: db,
	}

	params := postgres.InsertTaskParams{
		UserID:      uuid.MustParse("e05fa11d-eec3-4fba-b223-d6516800a047"),
		Description: "test description",
		Status:      1,
	}

	_, err = rp.InsertNewTask(context.Background(), params)
	if err != nil {
		log.Println(err)
		t.Fail()
	}
}
