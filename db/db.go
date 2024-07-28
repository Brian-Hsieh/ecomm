package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewSQLDatabase(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalln("Unable to connect to database")
	}

	return db, err
}
