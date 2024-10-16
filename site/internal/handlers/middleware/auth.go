package middleware

import (
	"net/http"
	"rapidart/internal/auth"
)

// Checks if the user is authenticated, redirects to login if not
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check auth and redirect if bad
		cookie, err := r.Cookie("session-token")
		if err == nil {
			_, err := auth.GetSession(cookie.Value)
			if err == nil {
				// Is authenticated!
				next.ServeHTTP(w, r) // Pass request to next handler
				return
			}
		}

		// Not authenticated! Redirect to login.
		http.Redirect(w, r, "/login/", http.StatusTemporaryRedirect)
	})
}
