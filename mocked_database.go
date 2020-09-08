package main

// A global variable incremented for providing a unique ID to each book
var count int

// A slice that will contain the books
var books []Book

func _get_all_books() []Book {
    return books
}

func _get_book(id int) Book {
	for _, book := range books {
		if book.ID == id {
		    return book
		}
	}
	return Book{}
}

func _insert_book(book Book) Book {
    book.ID = count
	books = append(books, book)
	count ++
	return book
}

func _update_book(tempBook Book) Book {
	for index, book := range books {
		if book.ID == tempBook.ID {
		    books[index] = tempBook
		    return tempBook
		}
	}
	return Book{}
}

func _delete_book(id int) {
	for index, book := range books {
		if book.ID == id {
            books = append(books[:index], books[index+1:]...)
		    return
		}
	}
}
