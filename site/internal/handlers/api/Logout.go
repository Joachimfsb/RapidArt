package api

import (
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/util"
)

// ////////////// HANDLER /////////////// //

// AuthLogout handler. This function routes the different REST methods to other handlers.
func AuthLogout(w http.ResponseWriter, r *http.Request) {

	// Get session cookie
	cookie, err := r.Cookie("session-token")
	if err != nil {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Perform logout
	auth.Logout(cookie.Value)

	// Set session cookie to none
	http.SetCookie(w, &http.Cookie{
		Name:  "session-token",
		Value: "",
		Path:  "/",

		HttpOnly: true, // Don't allow javascript to access cookie
	})

	w.WriteHeader(http.StatusNoContent)
}
