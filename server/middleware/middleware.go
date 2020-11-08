package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/a-edwaar/jwt/server/auth"
	"github.com/a-edwaar/jwt/server/models"
	"github.com/dgrijalva/jwt-go"
)

// Auth to check access token in header is valid
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// We can obtain the access token from the auth header
		bearerToken := r.Header.Get("Authorization")
		splitToken := strings.Split(bearerToken, "Bearer ")
		if len(splitToken) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tknStr := splitToken[1]
		// Validate token
		claims := &auth.Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return auth.JWTKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Set username in the context
		ctx := context.WithValue(r.Context(), models.ContextKey("username"), tkn.Claims.(*auth.Claims).Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// Auth to check refresh token is valid
func RefreshAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get cookie
		c, err := r.Cookie("refresh")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Validate token
		tknStr := c.Value
		claims := &auth.RefreshClaims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return auth.JWTKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Set id in the context
		ctx := context.WithValue(r.Context(), models.ContextKey("userID"), tkn.Claims.(*auth.RefreshClaims).UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
