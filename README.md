Abstract
======
    Author:    José Pereira

Install Requirements
----------
[Download Golang](https://golang.org/dl/)

Install Dependencies
----------
	go get -u github.com/dgrijalva/jwt-go
	go get -u github.com/gorilla/mux
	go get -u github.com/jinzhu/gorm
	go get -u github.com/joho/godotenv
	go get -u golang.org/x/crypto

Run
----------
    go run .

Test It With Python
----------
    import requests

    base_url = "http://localhost:8000/"

    requests.get(base_url).content
    requests.post(base_url + "register", json={"email": "test", "password": "test"}).json()
    token = requests.post(base_url + "login", json={"email": "test", "password": "test"}).json()["token"]

    requests.get(base_url + "api/user", headers={"x-access-token": token}).json()
    requests.get(base_url + "api/user/1", headers={"x-access-token": token}).json()
    requests.put(base_url + "api/user/1", json={"name": "test"}, headers={"x-access-token": token}).json()
    requests.delete(base_url + "api/user/1", headers={"x-access-token": token}).json()

    requests.post(base_url + "api/books", json={"title":"test","author":"test", "isbn": "1234"}, headers={"x-access-token": token}).json()
    requests.get(base_url + "api/books", headers={"x-access-token": token}).json()
    requests.get(base_url + "api/books/1", headers={"x-access-token": token}).json()
    requests.put(base_url + "api/books/1", json={"title":"test2","author":"test2", "isbn": "1234"}, headers={"x-access-token": token}).json()
    requests.delete(base_url + "api/books/1", headers={"x-access-token": token}).json()
