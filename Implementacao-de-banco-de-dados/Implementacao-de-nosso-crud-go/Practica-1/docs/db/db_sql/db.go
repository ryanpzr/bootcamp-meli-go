package db_sql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type Database struct{}

func (db *Database) OpenConnection() (*sql.DB, error) {
	conn, err := sql.Open("mysql", os.Getenv("DATABASE_DSN"))
	if err != nil {
		log.Fatal(err)
	}

	return conn, err
}
