package api

import (
	"log"
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/util"
)

// ////////////// HANDLER /////////////// //

// Login handler. This function routes the different REST methods to other handlers.
func Login(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		loginPost(w, r)
	default: //Error message if GET method is not used
		http.Error(w, "This method is not supported.", http.StatusNotImplemented)
	}
}

// Internal post handler for this route
func loginPost(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() // Neccessary?

	// Get data from form
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	// If either is missing, return error
	if username == "" || password == "" {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	token, err := auth.Login(username, password)
	if err != nil {
		log.Println(err)
		util.HttpReturnError(http.StatusUnauthorized, w)
		return
	}

	cookie := &http.Cookie{
		Name:  "session-token",
		Value: token,
		Path:  "/",

		HttpOnly: true, // Don't allow javascript to access cookie
	}

	http.SetCookie(w, cookie)

	// Redirect to front page
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
