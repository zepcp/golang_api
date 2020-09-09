package api

import (
    "net/http"
    "encoding/json"
    "strconv"
    "github.com/gorilla/mux"
	"golang_api/lib"
	"golang_api/db"
	// db "golang_api/db/mocked_db"
)

// Give all the books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(db.GetAllBooks())
}

// Give a book with some ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

    index, err := strconv.Atoi(vars["id"])
    if err == nil {
        json.NewEncoder(w).Encode(db.GetBook(index))
        return
    }

	json.NewEncoder(w).Encode(&lib.Book{})
}

// Adds a new Book
func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book lib.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

    book = db.InsertBook(book)
	json.NewEncoder(w).Encode(book)
}

// Updates a book with some ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	var tempBook lib.Book
    _ = json.NewDecoder(r.Body).Decode(&tempBook)

    index, err := strconv.Atoi(vars["id"])
    if err == nil {
        tempBook.ID = index
        db.UpdateBook(tempBook)
        return
    }

	json.NewEncoder(w).Encode(&lib.Book{})
}

// Deletes the book with some ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

    index, err := strconv.Atoi(vars["id"])
    if err == nil {
        db.DeleteBook(index)
        return
    }

	json.NewEncoder(w).Encode(&lib.Book{})
}
