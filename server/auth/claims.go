package auth

import "github.com/dgrijalva/jwt-go"

// struct for access token
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// struct for refresh token
type RefreshClaims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}
