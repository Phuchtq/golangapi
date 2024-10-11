package main

import (
	"log"
	"v3/controllers"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file - ", err)
	}
	//--------------------------------------
	controllers.InitializeAPIRoutes()
}
