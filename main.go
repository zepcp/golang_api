package main

import (
	"golang_api/restapi"
	"golang_api/restapi/operations"
	"log"
	"os"
	"strconv"

	"github.com/go-openapi/loads"
	"github.com/joho/godotenv"
)

func main() {
	//Load env variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("apiPort"))

	if err != nil {
		log.Fatal("Error loading apiPort from .env file")
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewGoLangAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.EnabledListeners = []string{"http"}
	server.Port = port

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
