package web

import (
	"log"
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/models"
	"rapidart/internal/user"
	"rapidart/internal/util"
)

// /////////////// TEMPLATE MODEL ///////////////////
type profileTemplateModel struct {
	IsSelf bool // Is this the logged in users account?
	User   models.User
}

// /////////////// HANDLER //////////////////
func Profile(w http.ResponseWriter, r *http.Request) {

	var isSelf bool
	var u models.User

	username := r.PathValue("username")
	if username == "" {
		// Verdi er tom, hent innlogget bruker
		isSelf = true

		// Get session cookie
		cookie, err := r.Cookie("session-token")
		if err != nil {
			util.HttpReturnError(http.StatusUnauthorized, w)
			return
		}

		// Get currently logged in user
		u, err = auth.GetLoggedInUser(cookie.Value)
		if err != nil {
			util.HttpReturnError(http.StatusUnauthorized, w)
			return
		}
	} else {
		// Brukernavn spesifiert, hent brukerinfo
		isSelf = false

		var err error
		u, err = user.GetUserByUsername(username)
		if err != nil {
			util.HttpReturnError(http.StatusBadRequest, w)
			return
		}
	}

	// Create model
	model := profileTemplateModel{
		IsSelf: isSelf,
		User:   u,
	}

	err := util.HttpServeTemplate("profile.tmpl", model, w)
	if err != nil {
		log.Println(err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
