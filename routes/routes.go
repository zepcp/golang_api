package routes

import (
	"auth/controllers"
	"auth/utils/auth"
	"net/http"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/register", controllers.AddUser).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	s := r.PathPrefix("/api").Subrouter()
	s.Use(auth.JwtVerify)
	s.HandleFunc("/users", controllers.FetchUsers).Methods("GET")
	s.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	s.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	s.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	s.HandleFunc("/books", controllers.AddBook).Methods("POST")
	s.HandleFunc("/books", controllers.FetchBooks).Methods("GET")
	s.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	s.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	s.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

	return r
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
