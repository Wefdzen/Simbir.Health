package main

import (
	"log"

	"wefdzen/api/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() //by default .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := router.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
