package api

import (
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/util"
)

// ////////////// HANDLER /////////////// //

// Logout handler. This function routes the different REST methods to other handlers.
func Logout(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		logoutPost(w, r)
	default: //Error message if GET method is not used
		http.Error(w, "This method is not supported.", http.StatusNotImplemented)
	}
}

// Internal post handler for this route
func logoutPost(w http.ResponseWriter, r *http.Request) {

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
