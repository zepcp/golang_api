Abstract
======
    Author:    José Pereira

Install Requirements
----------
[Download Golang](https://golang.org/dl/)

Install dependencies

    go get -u github.com/gorilla/mux
    go get -u github.com/lib/pq

Run
----------
    go run .

Test It
----------
    curl -X GET http://localhost:8000/
    curl -X GET http://localhost:8000/healthCheck
    curl -X GET http://localhost:8000/api/books
    curl -X POST http://localhost:8000/api/books --data '{"title":"test","author":{"firstname": "name", "lastname": "last"}}'
    curl -X PUT http://localhost:8000/api/books/0 --data '{"title":"updated title","author":{"firstname": "name", "lastname": "last"}}'
    curl -X GET http://localhost:8000/api/books/0
    curl -X DELETE http://localhost:8000/api/books/0
