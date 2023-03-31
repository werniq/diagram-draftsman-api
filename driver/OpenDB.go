package driver

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func OpenDB() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
