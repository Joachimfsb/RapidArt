package api

import (
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/util"
)

// /////////// MODEL //////////// //
type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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
	var loginData loginRequest
	err := util.JsonDecode(r.Body, &loginData)
	if err != nil {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// If either is missing, return error
	if loginData.Username == "" || loginData.Password == "" {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	token, wrongUser, wrongPass, err := auth.Login(loginData.Username, loginData.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "text/plain")
		if wrongUser {
			w.Write([]byte("bad-user"))
		} else if wrongPass {
			w.Write([]byte("bad-pass"))
		}
		return
	}

	cookie := &http.Cookie{
		Name:  "session-token",
		Value: token,
		Path:  "/",

		HttpOnly: true, // Don't allow javascript to access cookie
	}

	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusNoContent)
}
