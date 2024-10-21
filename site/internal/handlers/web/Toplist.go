package web

import (
	"log"
	"net/http"
	"rapidart/internal/basismanager"
	"rapidart/internal/models"
	post "rapidart/internal/post/like"
	follow "rapidart/internal/user/follow"
	"rapidart/internal/util"
	"time"
)

type ToplistPageData struct {
	BasisCanvases []models.BasisCanvas
	TopPosts      []models.PostExtended
	TopUsers      []models.UserExtended
}

func Toplist(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		toplistGetRequest(w, r)
	default:
		http.Error(w, "this method is not supported", http.StatusNotImplemented)
	}
}

func toplistGetRequest(w http.ResponseWriter, r *http.Request) {
	// Get current time
	currentTime := time.Now()

	// Fetch basis canvases
	canvases, err := basismanager.GetBasisCanvasesByDateTime(currentTime)
	if err != nil {
		log.Println("Error fetching basis canvases:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Fetch top 10 liked posts
	topPosts, err := post.GetTopLikedPosts(10)
	if err != nil {
		log.Println("Error fetching top liked posts:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Fetch top 10 followed users
	topUsers, err := follow.GetTopFollowedUsers(10)
	if err != nil {
		log.Println("Error fetching top followed users:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Prepare the data to send to the template
	pageData := ToplistPageData{
		BasisCanvases: canvases,
		TopPosts:      topPosts,
		TopUsers:      topUsers,
	}

	// Render toplist.tmpl with template for basis canvases, top posts, and top users
	err = util.HttpServeTemplate("toplist.tmpl", pageData, w)
	if err != nil {
		log.Println("Error serving template:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
