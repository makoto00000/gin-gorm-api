package main

import (
	"log"
	"app/api/config"
	"app/api/routes"
	"app/api/services"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("Could not connect to the database", err)
	}
	services.InitDB(db)

	router := routes.SetupRouter()
	router.Run(":3000")
}