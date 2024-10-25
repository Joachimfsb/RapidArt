package web

import (
	"log"
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/models"
	"rapidart/internal/util"
)

// /////////////// TEMPLATE MODEL ///////////////////
type profileTemplateModel struct {
	User models.User
}

///////////////// HANDLER //////////////////

func Profile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		profileGetRequest(w, r)
	default:
		http.Error(w, "this method is not supported", http.StatusNotImplemented)
	}
}

func profileGetRequest(w http.ResponseWriter, r *http.Request) {

	// Get session cookie
	cookie, err := r.Cookie("session-token")
	if err != nil {
		util.HttpReturnError(http.StatusUnauthorized, w)
		return
	}

	// Get currently logged in user
	u, err := auth.GetLoggedInUser(cookie.Value)
	if err != nil {
		util.HttpReturnError(http.StatusUnauthorized, w)
		return
	}

	// Create model
	model := profileTemplateModel{
		User: u,
	}

	err = util.HttpServeTemplate("profile.tmpl", model, w)
	if err != nil {
		log.Println(err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
