package main

import (
	"auth/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//Load env variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Handle routes
	http.Handle("/", routes.Handlers())

	// serve
	log.Fatal(http.ListenAndServe(":"+os.Getenv("apiPort"), nil))
}
