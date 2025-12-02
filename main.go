package main

import (
	"log"

	"github.com/raghavendrah25/backend/server"
)

func main() {
	myServer := server.New(server.Config{})

	err := myServer.Run()
	if err != nil {
		log.Fatal(err)
	}
}
