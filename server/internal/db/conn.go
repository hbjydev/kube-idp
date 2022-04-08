package db

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"
)

// Db represents a connection to the database.
var Db *sql.DB

// Connect opens a connection to the database and stores it for the
// duration of the application execution.
func Connect() error {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return errors.New("no DATABASE_URL specified")
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	Db = db
	return nil
}
