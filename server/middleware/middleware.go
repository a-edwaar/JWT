package middleware

import (
	"context"
	"net/http"

	"github.com/a-edwaar/jwt/server/auth"
	"github.com/a-edwaar/jwt/server/models"
	"github.com/dgrijalva/jwt-go"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// We can obtain the session token from the requests cookies, which come with every request
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Validate token
		tknStr := c.Value
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
