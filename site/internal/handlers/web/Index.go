package web

import (
	"log"
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/basismanager"
	"rapidart/internal/models"
	"rapidart/internal/post"
	"rapidart/internal/util"
	"time"
)

type IndexPageData struct {
	Title              string
	BasisCanvases      []models.BasisCanvas
	RecentFollowsPosts []models.PostExtended
	RecentPosts        []models.PostExtended
}

func Index(w http.ResponseWriter, r *http.Request) {

	// Get logged in user
	session, err := auth.GetSession(util.GetSessionTokenFromCookie(r))
	if err != nil {
		util.HttpReturnError(http.StatusUnauthorized, w)
		return
	}

	// Get current time
	currentTime := time.Now()

	// Gets list of basis canvases based on current time
	canvases, err := basismanager.GetBasisCanvasesByDateTime(currentTime)
	if err != nil {
		log.Println("Error fetching basis canvases:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Get recent follows posts
	recentFollowsPosts, err := post.GetRecentFollowsPosts(session.UserId, 10)
	if err != nil {
		log.Println("Error fetching feed:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Get recent posts
	recentPosts, err := post.GetRecentPosts(10)
	if err != nil {
		log.Println("Error fetching feed:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Prepare the data to send to the template
	pageData := IndexPageData{
		BasisCanvases:      canvases,
		RecentFollowsPosts: recentFollowsPosts,
		RecentPosts:        recentPosts,
	}

	// Renders index.tmpl with template for basis canvases
	err = util.HttpServeTemplate("index.tmpl", pageData, w)
	if err != nil {
		log.Println("Error serving template:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
