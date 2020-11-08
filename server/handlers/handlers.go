package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/a-edwaar/jwt/server/auth"
	"github.com/a-edwaar/jwt/server/models"
)

// for testing - in reality will check db
var users = map[string]models.User{
	"1": models.User{
		Username: "user1",
		Password: "pass1",
	},
	"2": models.User{
		Username: "user2",
		Password: "pass2",
	},
}

func Login(w http.ResponseWriter, req *http.Request) {
	// Get creds from request
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Get the expected password from our in memory map
	userFound := false
	for _, storedUser := range users {
		if storedUser.Username == user.Username && storedUser.Password == user.Password {
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Generate access + refresh token
	accessToken, refreshCookie, err := auth.GenerateTokenPair(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Set refresh token cookie
	http.SetCookie(w, refreshCookie)
	//Write access token in response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(accessToken)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	// Get id from context - middleware sets from refresh cookie val
	userID := r.Context().Value(models.ContextKey("userID"))
	if userID == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Get user from db for id in token
	userToLookup := users[userID.(string)]
	// Generate access + refresh token
	accessToken, refreshCookie, err := auth.GenerateTokenPair(userToLookup)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Set refresh token cookie
	http.SetCookie(w, refreshCookie)
	//Write access token in response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(accessToken)
}

func Private(w http.ResponseWriter, r *http.Request) {
	// Get username from context - middleware sets from refresh cookie val
	username := r.Context().Value(models.ContextKey("username"))
	if username == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Printf("hello im in the restricted section for %s", username.(string))
}
