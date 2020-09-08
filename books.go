package main

import (
    "net/http"
    "encoding/json"
    "strconv"
    "github.com/gorilla/mux"
)

// Book Struct
type Book struct {
	ID     int     `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Give all the books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(get_all_books())
}

// Give a book with some ID
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

    index, err := strconv.Atoi(vars["id"])
    if err == nil {
        json.NewEncoder(w).Encode(get_book(index))
        return
    }

	json.NewEncoder(w).Encode(&Book{})
}

// Adds a new Book
func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

    book = insert_book(book)
	json.NewEncoder(w).Encode(book)
}

// Updates a book with some ID
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	var tempBook Book
    _ = json.NewDecoder(r.Body).Decode(&tempBook)

    index, err := strconv.Atoi(vars["id"])
    if err == nil {
        tempBook.ID = index
        update_book(tempBook)
        return
    }

	json.NewEncoder(w).Encode(&Book{})
}

// Deletes the book with some ID
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

    index, err := strconv.Atoi(vars["id"])
    if err == nil {
        delete_book(index)
        return
    }

	json.NewEncoder(w).Encode(&Book{})
}
