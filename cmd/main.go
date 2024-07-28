package main

import (
	"log"

	"github.com/Brian-Hsieh/ecomm/cmd/ecomm"
)

func main() {
	server := ecomm.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatalln("Unable to run the API server")
	}
}
