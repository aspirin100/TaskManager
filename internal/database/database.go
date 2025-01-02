package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/aspirin100/TaskManager/internal/migrations"

	"github.com/google/uuid"
	"github.com/pressly/goose/v3"
)

type PostgresRepo struct {
	DB *sql.DB
}

type InsertTaskParams struct {
	UserID      uuid.UUID
	Type        string
	Name        string
	Description string
	Status      uint8
}

func UpDatabase(driver, DSN string) (*sql.DB, error) {
	db, err := sql.Open(driver, DSN)
	if err != nil {
		return nil, fmt.Errorf("open database error: %w", err)
	}

	goose.SetBaseFS(migrations.Migrations)

	err = goose.Up(db, ".")
	if err != nil {
		return nil, fmt.Errorf("migrations up error: %w", err)
	}

	return db, nil
}

func (pg *PostgresRepo) InsertNewTask(ctx context.Context, params InsertTaskParams) (uuid.UUID, error) {

	taskID := uuid.New()

	_, err := pg.DB.QueryContext(ctx, InsertTaskQuery, taskID,
		params.UserID,
		params.Type,
		params.Name,
		params.Description,
		params.Status,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("insert task query executing error: %w", err)
	}

	return taskID, nil
}

func (pg *PostgresRepo) DeleteTask(ctx context.Context, taskID uuid.UUID) error {

	_, err := pg.DB.QueryContext(ctx, DeleteTaskQuery, taskID)
	if err != nil {
		return fmt.Errorf("delete task query error: %w", err)
	}

	return nil
}

const (
	InsertTaskQuery = `insert into tasks(taskID, userID, type, name, description, status) values ($1, $2, $3, $4, $5, $6)`
	DeleteTaskQuery = `delete from tasks where taskID = $1`
)
