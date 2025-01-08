package tasksRepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/aspirin100/TaskManager/internal/migrations"
	"github.com/aspirin100/TaskManager/internal/tasks"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type PostgresRepo struct {
	DB *sql.DB
}

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrUserNotFound = errors.New("user doesn't exist")
)

const (
	ForeignKeyViolationCode = "23503"
)

func New(DSN string) (*PostgresRepo, error) {
	db, err := upDatabase("postgres", DSN)
	if err != nil {
		return &PostgresRepo{}, fmt.Errorf("posgtres repository creation error: %w", err)
	}

	return &PostgresRepo{
		DB: db,
	}, nil
}

func upDatabase(driver, DSN string) (*sql.DB, error) {
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

func (pg *PostgresRepo) CreateTask(ctx context.Context, params tasks.CreateTaskRequest) (uuid.UUID, error) {
	taskID := uuid.New()

	_, err := pg.DB.ExecContext(ctx, InsertTaskQuery, taskID,
		params.UserID,
		params.Type,
		params.Name,
		params.Description,
		params.Status,
	)
	if err != nil {
		var pgErr *pq.Error

		switch {
		case errors.As(err, &pgErr) && pgErr.Code == ForeignKeyViolationCode:
			return uuid.Nil, ErrUserNotFound
		default:
			return uuid.Nil, fmt.Errorf("insert task query error: %w", err)
		}

	}

	return taskID, nil
}

func (pg *PostgresRepo) DeleteTask(ctx context.Context, params tasks.CommonTaskRequest) error {

	res, err := pg.DB.ExecContext(ctx, DeleteTaskQuery, params.TaskID, params.UserID)
	if err != nil {
		return fmt.Errorf("delete task query error: %w", err)

	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("affected rows checking error: %w", err)
	}

	if rows == 0 {
		return ErrTaskNotFound
	}

	return nil
}

func (pg *PostgresRepo) UpdateTask(ctx context.Context, params tasks.UpdateTaskRequest) (uuid.UUID, error) {

	res, err := pg.DB.ExecContext(ctx, UpdateTaskQuery,
		params.TaskID,
		params.Type,
		params.Name,
		params.Description,
		params.Status,
		params.UserID,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("update task query error: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return uuid.Nil, fmt.Errorf("affected rows checking error: %w", err)
	}

	if rows == 0 {
		return uuid.Nil, ErrTaskNotFound
	}

	return params.TaskID, nil
}

func (pg *PostgresRepo) GetTask(ctx context.Context, params tasks.CommonTaskRequest) (tasks.Task, error) {
	row, err := pg.DB.QueryContext(ctx, GetTaskQuery, params.TaskID, params.UserID)
	if err != nil {
		return tasks.Task{}, fmt.Errorf("get task query error: %w", err)
	}

	res := tasks.Task{}
	userID := uuid.Nil

	if !row.Next() {
		return tasks.Task{}, ErrTaskNotFound
	}

	err = row.Scan(
		&res.TaskID,
		&userID,
		&res.Type,
		&res.Name,
		&res.Description,
		&res.Status,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return tasks.Task{}, fmt.Errorf("columns scanning query error: %w", err)
	}

	return res, nil
}

func (pg *PostgresRepo) CheckUserExists(ctx context.Context, userID uuid.UUID) error {
	res, err := pg.DB.ExecContext(ctx, CheckUserQuery, userID)
	if err != nil {
		return fmt.Errorf("user check query error: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("user check rows affected checking error: %w", err)
	}

	if rows == 0 {
		return ErrUserNotFound
	}

	return nil
}

const (
	InsertTaskQuery = `insert into tasks(taskID, userID, type, name, description, status) values ($1, $2, $3, $4, $5, $6)`
	DeleteTaskQuery = `delete from tasks where taskID = $1 and userID = $2`
	UpdateTaskQuery = `update tasks set type = $2, name = $3, description = $4, status = $5, updatedAt = now() where taskID = $1 and userID = $6`
	GetTaskQuery    = `select * from tasks where taskID = $1 and userID = $2`
	CheckUserQuery  = `select * from users where id = $1`
)
