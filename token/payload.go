package token

import "github.com/golang-jwt/jwt"

type Payload struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
