package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/aspirin100/TaskMaster/migrations"

	"github.com/pressly/goose/v3"
)

type PostgresRepo struct {
	db *sql.DB
}

func UpDatabase(driver, DSN string) error {
	db, err := sql.Open(driver, DSN)
	if err != nil {
		return fmt.Errorf("open database error: %w", err)
	}

	goose.SetBaseFS(migrations.Migrations)

	err = goose.Up(db, ".")
	if err != nil {
		return fmt.Errorf("migrations up error: %w", err)
	}

	return nil
}
