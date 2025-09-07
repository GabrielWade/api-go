package main

import (
	"api/app/infrastructure/config"
	"api/app/interfaces/controller"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.NewDatabaseConnection()
}

func main() {
	router := controller.NewRouter()
	router.Run()
}
