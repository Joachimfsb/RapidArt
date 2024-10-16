package api

import (
	"net/http"
	"rapidart/internal/models"
	"rapidart/internal/user"
	"rapidart/internal/util"
)

// ////////////// HANDLER /////////////// //

// Registration handler. This function routes the different REST methods to other handlers.
func Register(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		registerPost(w, r)
	default: //Error message if GET method is not used
		http.Error(w, "This method is not supported.", http.StatusNotImplemented)
	}
}

// Internal post handler for this route
func registerPost(w http.ResponseWriter, r *http.Request) {

	// Parse request data
	var registerData models.RegisterUser
	err := util.JsonDecode(r.Body, &registerData)
	if err != nil {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Try to create user
	err = user.CreateUser(registerData)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		// Safe to return error code because CreateUser specifies it
		w.Write([]byte(err.Error())) // Lets javascript know whats wrong

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
