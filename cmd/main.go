package main

import (
	"log"

	"github.com/Brian-Hsieh/ecomm/cmd/ecomm"
	"github.com/Brian-Hsieh/ecomm/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewSQLDatabase(mysql.Config{})

	server := ecomm.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatalln("Unable to run the API server")
	}
}
