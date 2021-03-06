package repository

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/pkg/errors"
)

func NewDB() (*sql.DB, error) {
	dsn := os.Getenv("DSN")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sql.Open")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "db.Ping() failed")
	}

	return db, nil
}
