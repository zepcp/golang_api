package auth

import (
	"auth/models"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

//Exception type - struct
type Exception models.Exception

// JwtVerify func - Middleware
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var accessToken = r.Header.Get("x-access-token")

		accessToken = strings.TrimSpace(accessToken)

		if accessToken == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"})
			return
		}
		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(accessToken, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
