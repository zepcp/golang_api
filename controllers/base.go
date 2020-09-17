package controllers

import (
	"auth/utils"
	"net/http"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

var db = utils.ConnectDB()

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API live and kicking"))
}