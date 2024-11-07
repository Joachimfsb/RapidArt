package api

import (
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/glob"
	"rapidart/internal/util"
	"time"
)

// /////////// MODEL //////////// //
type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ////////////// HANDLER /////////////// //

// AuthLogin handler. This function routes the different REST methods to other handlers.
func AuthLogin(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		authLoginPost(w, r)
	default: //Error message if GET method is not used
		http.Error(w, "This method is not supported.", http.StatusNotImplemented)
	}
}

// Internal post handler for this route
func authLoginPost(w http.ResponseWriter, r *http.Request) {

	// Parse request data
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

	// Perform login and check status
	token, wrongUser, wrongPass, err := auth.Login(loginData.Username, loginData.Password, r.RemoteAddr, util.UserAgentToBrowser(r.UserAgent()))
	if err != nil {
		// Could not log in
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "text/plain")
		if wrongUser || wrongPass { // Problem was that the user does not exist or the password is wrong
			w.Write([]byte("bad-creds"))
		}
		return
	}

	// Return token to user
	cookie := &http.Cookie{
		Name:    "session-token",
		Value:   token,
		Path:    "/",
		Expires: time.Now().AddDate(0, 0, glob.SessionExpirationDays),

		HttpOnly: true, // Don't allow javascript to access cookie
	}

	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusNoContent)
}
