package main

import (
	"log"
	"net/http"

	"github.com/a-edwaar/jwt/server/handlers"
	"github.com/a-edwaar/jwt/server/middleware"
)

func main() {
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/refresh", middleware.RefreshAuth(handlers.Refresh))
	http.HandleFunc("/private", middleware.Auth(handlers.Private))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
