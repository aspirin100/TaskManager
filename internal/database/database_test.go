package database_test

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/aspirin100/TaskManager/internal/database"
	"github.com/google/uuid"
)

func TestInsertNewTaskFail(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/task-manager?sslmode=disable")
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	rp := database.PostgresRepo{
		DB: db,
	}

	params := database.InsertTaskParams{
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

	rp := database.PostgresRepo{
		DB: db,
	}

	params := database.InsertTaskParams{
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

func TestDeleteTask(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/task-manager?sslmode=disable")
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	rp := database.PostgresRepo{
		DB: db,
	}

	err = rp.DeleteTask(context.Background(), uuid.MustParse("f182607f-6b8f-45a9-ad47-b6d38013c827"))
	if err != nil {
		log.Println(err)
		t.Fail()
	}
}
