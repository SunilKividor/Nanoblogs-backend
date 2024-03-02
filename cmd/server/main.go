package main

import (
	"log"

	"github.com/SunilKividor/internal/repository/postgresql"
	"github.com/SunilKividor/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	postgresql.ConnectDB()
	err = server.StartServer()
	if err != nil {
		log.Fatal("Error starting the server")
	}
}
