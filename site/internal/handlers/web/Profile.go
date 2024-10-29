package web

import (
	"log"
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/models"
	"rapidart/internal/post"
	"rapidart/internal/user"
	"rapidart/internal/util"
	"slices"
	"strings"
)

// /////////////// TEMPLATE MODEL ///////////////////
type profileTemplateModel struct {
	IsSelf     bool // Is this the logged in users account?
	IsFollower bool // Is the logged in user a follower of the profile beeing viewed
	User       models.User
	PostList   []models.PostExtended
	Stats      models.UserStats
}

// /////////////// HANDLER //////////////////
func Profile(w http.ResponseWriter, r *http.Request) {

	//// Get currently logged in user ////
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

	// Check if asked upon user is selv or not.
	var isSelf bool
	var u models.User

	username := strings.ToLower(r.PathValue("username"))
	if username == "" || username == loggedInUser.Username {
		// Request is asking for currently logged in users profile
		isSelf = true
		u = loggedInUser

	} else {
		// Request is asking for another user, fetch their data...
		isSelf = false

		var err error
		u, err = user.GetUserByUsername(username)
		if err != nil {
			util.HttpReturnError(http.StatusBadRequest, w)
			return
		}
	}

	// -- Get user posts -- //
	p, err := post.GetRecentPostsByUser(u.UserId, 10)
	if err != nil {
		log.Println(err.Error())
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// -- Get user stats -- //
	stats, err := user.GetUserStats(u.UserId)
	if err != nil {
		log.Println(err.Error())
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// -- Is follower -- //
	isFollower := false
	if !isSelf && slices.Contains(stats.Followers, loggedInUser.UserId) {
		isFollower = true
	}

	// Create model
	model := profileTemplateModel{
		IsSelf:     isSelf,
		IsFollower: isFollower,
		User:       u,
		PostList:   p,
		Stats:      stats,
	}

	err = util.HttpServeTemplate("profile.tmpl", model, w)
	if err != nil {
		log.Println(err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
