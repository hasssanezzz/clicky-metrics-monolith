package domain

import "github.com/golang-jwt/jwt/v4"

type ErrorResponse struct {
	Message string `json:"message"`
}

type JwtCustomClaims struct {
	Username string `json:"name"`
	ID       int    `json:"id"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}
