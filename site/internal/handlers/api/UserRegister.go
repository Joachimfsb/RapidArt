package api

import (
	"html"
	"log"
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/models"
	"rapidart/internal/user"
	"rapidart/internal/util"
)

// ////////////// HANDLER /////////////// //

// User registration handler. This function routes the different REST methods to other handlers.
func UserRegister(w http.ResponseWriter, r *http.Request) {

	// Parse request data
	var registerData models.RegisterUser
	err := util.JsonDecode(r.Body, &registerData)
	if err != nil {
		log.Println(err)
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Parameter specifies that the requester only wishes to check if email and username is available
	if r.URL.Query().Has("check_email_username") {
		UserRegisterCheckEmailUsername(registerData, w, r)
		return
	}

	// HTML sanitize input
	registerData.Email = html.EscapeString(registerData.Email)
	registerData.Username = html.EscapeString(registerData.Username)
	registerData.Displayname = html.EscapeString(registerData.Displayname)

	// Try to create user
	err = user.CreateUser(registerData)
	if err != nil {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Perform login and check status
	token, _, _, err := auth.Login(registerData.Username, registerData.Password)
	if err != nil {
		// Could not log in (should not happen)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Return token to user
	cookie := &http.Cookie{
		Name:  "session-token",
		Value: token,
		Path:  "/",

		HttpOnly: true, // Don't allow javascript to access cookie
	}

	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusNoContent)
}

// User registration handler
func UserRegisterCheckEmailUsername(data models.RegisterUser, w http.ResponseWriter, r *http.Request) {

	if !user.CheckEmailAvailability(data.Email) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("email-exists")) // Lets javascript know whats wrong
		return
	}

	if !user.CheckUsernameAvailability(data.Username) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("username-exists")) // Lets javascript know whats wrong
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
