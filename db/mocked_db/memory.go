package mocked_db

import (
	"golang_api/lib"
)

// A global variable incremented for providing a unique ID to each book
var count int

// A slice that will contain the books
var books []lib.Book

func GetAllBooks() []lib.Book {
    return books
}

func GetBook(id int) lib.Book {
	for _, book := range books {
		if book.ID == id {
		    return book
		}
	}
	return lib.Book{}
}

func InsertBook(book lib.Book) lib.Book {
    book.ID = count
	books = append(books, book)
	count ++
	return book
}

func UpdateBook(tempBook lib.Book) lib.Book {
	for index, book := range books {
		if book.ID == tempBook.ID {
		    books[index] = tempBook
		    return tempBook
		}
	}
	return lib.Book{}
}

func DeleteBook(id int) {
	for index, book := range books {
		if book.ID == id {
            books = append(books[:index], books[index+1:]...)
		    return
		}
	}
}
