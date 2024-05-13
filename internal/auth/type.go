package auth

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	jwt.RegisteredClaims
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	ID        uint   `json:"id"`
}
