Abstract
======
    Author:    José Pereira

Install Requirements
----------
[Download Golang](https://golang.org/dl/)

Run
----------
    go run .

Test It With Python
----------
    import requests

    requests.get("http://localhost:8000/").content
    requests.post("http://localhost:8000/register", json={"email": "test", "password": "test"}).json()
    token = requests.post("http://localhost:8000/login", json={"email": "test", "password": "test"}).json()["token"]

    requests.get("http://localhost:8000/api/user", headers={"x-access-token": token}).json()
    requests.get("http://localhost:8000/api/user/1", headers={"x-access-token": token}).json()
    requests.put("http://localhost:8000/api/user/1", json={"name": "test"}, headers={"x-access-token": token}).json()
    requests.delete("http://localhost:8000/api/user/1", headers={"x-access-token": token}).json()

    requests.get("http://localhost:8000/api/books", headers={"x-access-token": token}).json()
    requests.get("http://localhost:8000/api/books/1", headers={"x-access-token": token}).json()
    requests.post("http://localhost:8000/api/books", json={"title":"test","author":"test", "isbn": "1234"}, headers={"x-access-token": token}).json()

    requests.put('http://localhost:8000/api/books/1', json={"title":"test2","author":"test2", "isbn": "1234"}, headers={"x-access-token": token}).json()
    requests.delete('http://localhost:8000/api/books/1', headers={"x-access-token": token}).json()
