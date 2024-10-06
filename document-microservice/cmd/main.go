package main

import (
	"log"

	"wefdzen/api/router"
)

func main() {
	r := router.SetupRouter()
	log.Fatal(r.Run(":8083"))
}
