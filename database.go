/*
CREATE TABLE book (
    id serial primary key,
    isbn varchar(20),
    title varchar(100),
    author json,
    ts timestamp without time zone default (now() at time zone 'utc')
);
*/
package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = ""
    dbname   = "postgres"
)

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func connect() *sql.DB{
    // connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    // open database
    db, e := sql.Open("postgres", psqlconn)
    CheckError(e)

    // check db
    e = db.Ping()
    CheckError(e)
    return db
}

func get_all_books() []Book{
    db := connect()
    rows, e := db.Query(`SELECT id, isbn, title, author->>'firstname', author->>'lastname' FROM book`)
    CheckError(e)

    var book Book
    var author Author
    var books []Book

    for rows.Next() {
        rows.Scan(&book.ID, &book.Isbn, &book.Title, &author.FirstName, &author.LastName)
        book.Author = &author
        books = append(books, book)
    }

    // close database
    defer db.Close()
    defer rows.Close()

    return books
}

func get_book(id int) Book{
    fmt.Println(id)
    db := connect()
    row := db.QueryRow(`SELECT id, isbn, title, author->>'firstname', author->>'lastname' FROM book WHERE id=$1`, id)

    var book Book
    var author Author
    row.Scan(&book.ID, &book.Isbn, &book.Title, &author.FirstName, &author.LastName)
    book.Author = &author

    // close database
    defer db.Close()

    return book
}

func insert_book(book Book) Book {
    db := connect()
    insertDynStmt := `INSERT INTO book(isbn, title, author) VALUES($1, $2, $3)`
    author, _ := json.Marshal(book.Author)
    _, e := db.Exec(insertDynStmt, book.Isbn, book.Title, author)
    CheckError(e)

    // close database
    defer db.Close()

    return book
}

func update_book(tempBook Book) Book {
    db := connect()
    updateStmt := `UPDATE book SET isbn=$1, title=$2, author=$3 WHERE id=$4`
    author, _ := json.Marshal(tempBook.Author)
    _, e := db.Exec(updateStmt, tempBook.Isbn, tempBook.Title, author, tempBook.ID)
    CheckError(e)

    // close database
    defer db.Close()

    return tempBook
}

func delete_book(id int) {
    db := connect()
    deleteStmt := `DELETE FROM book WHERE id=$1`
    _, e := db.Exec(deleteStmt, id)
    CheckError(e)

    // close database
    defer db.Close()
}
