package middleware

import (
	"log"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check auth and redirect if bad
		log.Println("Checking auth....")
		//http.Redirect(w, r, "/login/", http.StatusTemporaryRedirect)
		//
		next.ServeHTTP(w, r) // Pass request to next handler
	})
}
