package object

import "github.com/dgrijalva/jwt-go"

// CustomClaims represent JWT claims
type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// CustomClaimsRefresh represent JWT claims
type CustomClaimsRefresh struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
