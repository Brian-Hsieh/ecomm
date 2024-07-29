package main

import (
	"log"

	"github.com/Brian-Hsieh/ecomm/cmd/ecomm"
	"github.com/Brian-Hsieh/ecomm/config"
	"github.com/Brian-Hsieh/ecomm/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewSQLDatabase(mysql.Config{
		User:                 config.Env.DBUser,
		Passwd:               config.Env.DBPassword,
		Addr:                 config.Env.DBAddress,
		DBName:               config.Env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	server := ecomm.NewAPIServer(config.Env.ServerAddress, db)
	if err := server.Run(); err != nil {
		log.Fatalln("Unable to start the API server")
	}
}
