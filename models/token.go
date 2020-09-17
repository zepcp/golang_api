package models

import jwt "github.com/dgrijalva/jwt-go"

// Token struct
type Token struct {
	UserID uint
	Name   string
	Email  string
	*jwt.StandardClaims
}
