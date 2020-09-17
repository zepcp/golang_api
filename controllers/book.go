package controllers

import (
	"auth/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func FetchBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	db.Preload("auths").Find(&books)

	json.NewEncoder(w).Encode(books)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	params := mux.Vars(r)
	var id = params["id"]
	db.First(&book, id)
	json.NewDecoder(r.Body).Decode(book)
	db.Save(&book)
	json.NewEncoder(w).Encode(&book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var book models.Book
	db.First(&book, id)
	db.Delete(&book)
	json.NewEncoder(w).Encode("Book deleted")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var book models.Book
	db.First(&book, id)
	json.NewEncoder(w).Encode(&book)
}

func AddBook(w http.ResponseWriter, r *http.Request) {

	book := &models.Book{}
	json.NewDecoder(r.Body).Decode(book)

	addedBook := db.Create(book)
	var errMessage = addedBook.Error

	if addedBook.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(addedBook)
}
