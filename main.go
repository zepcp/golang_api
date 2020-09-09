package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"golang_api/api"
)

// Home screen
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Congrats. Run Successfully!</h1>")
}

// Check if system is up and running
func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Alive!</h1>")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", home)
	router.HandleFunc("/healthCheck", healthCheck)

	// Handling the endpoints
	router.HandleFunc("/api/books", api.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", api.GetBook).Methods("GET")
	router.HandleFunc("/api/books", api.AddBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", api.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", api.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
