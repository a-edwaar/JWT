package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/a-edwaar/jwt/server/models"
	"github.com/dgrijalva/jwt-go"
)

// Secret key to check signature
var JWTKey = []byte("my_secret_key")

func GenerateTokenPair(user models.User) (accessCookie *http.Cookie, refreshCookie *http.Cookie, err error) {
	// Generate access token
	accessExpirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpirationTime.Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	at, err := accessToken.SignedString(JWTKey)
	if err != nil {
		return nil, nil, fmt.Errorf("Access token err: %s", err.Error())
	}
	// Generate refresh token
	refreshExpirationTime := time.Now().Add(time.Hour * 24)
	refreshClaims := &RefreshClaims{
		UserID: "1",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	rt, err := refreshToken.SignedString(JWTKey)
	if err != nil {
		return nil, nil, fmt.Errorf("Refresh token err: %s", err.Error())
	}
	// Create cookies
	accessCookie = &http.Cookie{
		Name:     "token",
		Value:    at,
		Expires:  accessExpirationTime,
		HttpOnly: true,
		// Secure: true,
	}
	refreshCookie = &http.Cookie{
		Name:     "refresh",
		Value:    rt,
		Expires:  refreshExpirationTime,
		HttpOnly: true,
		// Secure: true,
	}
	return accessCookie, refreshCookie, nil
}
