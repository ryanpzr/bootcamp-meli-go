package repository

import (
	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
)

func init() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "12345",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "fantasy_products_test",
	}

	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}
