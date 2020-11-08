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

// Struct for returning access token in response
type AccessToken struct {
	Token string
}

func GenerateTokenPair(user models.User) (accessToken *AccessToken, refreshCookie *http.Cookie, err error) {
	// Generate access token
	accessExpirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpirationTime.Unix(),
		},
	}
	accessJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	at, err := accessJWT.SignedString(JWTKey)
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
	refreshJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refreshJWT.SignedString(JWTKey)
	if err != nil {
		return nil, nil, fmt.Errorf("Refresh token err: %s", err.Error())
	}
	// Create access token struct for response
	accessToken = &AccessToken{
		Token: at,
	}
	// Create refresh token cookie
	refreshCookie = &http.Cookie{
		Name:     "refresh",
		Value:    refreshToken,
		Expires:  refreshExpirationTime,
		HttpOnly: true,
		// Secure: true,
	}
	return accessToken, refreshCookie, nil
}
