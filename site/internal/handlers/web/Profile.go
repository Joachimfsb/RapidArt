package web

import (
	"log"
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/models"
	"rapidart/internal/post"
	"rapidart/internal/user"
	"rapidart/internal/util"
	"strings"
)

// /////////////// TEMPLATE MODEL ///////////////////
type profileTemplateModel struct {
	IsSelf   bool // Is this the logged in users account?
	User     models.User
	PostList []models.PostExtended
}

// /////////////// HANDLER //////////////////
func Profile(w http.ResponseWriter, r *http.Request) {

	// Get session cookie
	cookie, err := r.Cookie("session-token")
	if err != nil {
		util.HttpReturnError(http.StatusUnauthorized, w)
		return
	}

	// Get currently logged in user
	loggedInUser, err := auth.GetLoggedInUser(cookie.Value)
	if err != nil {
		util.HttpReturnError(http.StatusUnauthorized, w)
		return
	}

	var isSelf bool
	var u models.User

	username := strings.ToLower(r.PathValue("username"))
	if username == "" || username == loggedInUser.Username {
		// Verdi er tom, hent innlogget bruker
		isSelf = true
		u = loggedInUser

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

	// Get user posts
	p, err := post.GetRecentPostsByUser(u.UserId, 10)
	if err != nil {
		log.Println(err.Error())
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Create model
	model := profileTemplateModel{
		IsSelf:   isSelf,
		User:     u,
		PostList: p,
	}

	err = util.HttpServeTemplate("profile.tmpl", model, w)
	if err != nil {
		log.Println(err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
