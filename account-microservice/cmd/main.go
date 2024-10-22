package main

import (
	"log"
	"os"

	"wefdzen/api/router"
	"wefdzen/internal/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("/app/account-microservice/.env") //by default .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//при перезапуске docker слетает переменная
	if os.Getenv("test") == "" {
		database.InitDbTask()
	}
	os.Setenv("test", "initSucces")

	r := router.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
