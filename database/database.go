package database

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewSQLDatabase(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalln("Unable to create database client")
	}

	return db, err
}

func CheckConnection(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatalln("Unable to connect to database")
	}
	log.Println("Connected to database")
}
