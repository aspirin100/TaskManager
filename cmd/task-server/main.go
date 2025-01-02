package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/aspirin100/TaskMaster/migrations"

	"github.com/pressly/goose/v3"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/auth?sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}

	goose.SetBaseFS(migrations.Migrations)

	err = goose.Up(db, ".")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("migrations up")

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		println("ListenAndServe error")
		os.Exit(1)
	}

}

type Config struct {
	PostgresDSN string `envconfig:"TASK_SERVER_POSTGRES_DSN" default:"postgres://postgres:postgres@localhost:5432/auth?sslmode=disable"` //nolint:lll
	Hostname    string `envconfig:"TASK_SERVER_HOSTNAME" default:":8000"`
}
