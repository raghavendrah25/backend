package main

import (
	"log"

	"github.com/raghavendrah25/backend/server"
)

func main() {
	myServer := server.New(server.Config{Port: 8080})

	err := myServer.Run()
	if err != nil {
		log.Fatal(err)
	}
}
