package main

import (
	"github/thallesrangel/api-book-go/database"
	"github/thallesrangel/api-book-go/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
