package web

import (
	"log"
	"net/http"
	"rapidart/internal/models"
	"rapidart/internal/post"
	"rapidart/internal/user"
	"rapidart/internal/user/follow"
	"rapidart/internal/util"
	"strconv"
	"time"
)

type TopPostsData struct {
	TopList []models.PostExtended
}
type TopUsersData struct {
	TopList []models.UserExtended
}

func Top(w http.ResponseWriter, r *http.Request) {
	// --- Parse params --- //
	topType := r.PathValue("type")
	if topType == "" { // Missing post_id
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	var since *time.Time = nil
	var basisCanvasId *int = nil
	var metric string

	if topType == "users" {
		// Parse query params
		metric = r.URL.Query().Get("metric")
		if metric == "" {
			util.HttpReturnError(http.StatusBadRequest, w)
			return
		}

	} else if topType == "posts" {
		// Parse query params
		sinceStr := r.URL.Query().Get("since")
		if sinceStr != "" {
			sinceTime, err := time.Parse(time.RFC3339, sinceStr) // Convert to time
			if err != nil {
				util.HttpReturnError(http.StatusBadRequest, w)
				return
			}
			since = new(time.Time)
			*since = sinceTime
		}

		basisCanvasIdStr := r.URL.Query().Get("basiscanvas")
		if basisCanvasIdStr != "" {
			basisCanvasIdInt, err := strconv.Atoi(basisCanvasIdStr) // Convert to int
			if err != nil {
				util.HttpReturnError(http.StatusBadRequest, w)
				return
			}
			basisCanvasId = new(int)
			*basisCanvasId = basisCanvasIdInt
		}

	} else {
		// Incorrect spesification of type
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// --- Fetch data --- //

	if topType == "users" {

		var top []models.UserExtended

		if metric == "likes" {

			// Fetch top 10 followed users
			var err error
			top, err = user.GetMostLikedUsers(10)
			if err != nil {
				log.Println("Error fetching top followed users:", err)
				util.HttpReturnError(http.StatusInternalServerError, w)
				return
			}

		} else if metric == "followers" {

			// Fetch top 10 followed users
			var err error
			top, err = follow.GetTopFollowedUsers(10)
			if err != nil {
				log.Println("Error fetching top followed users:", err)
				util.HttpReturnError(http.StatusInternalServerError, w)
				return
			}

		} else {
			util.HttpReturnError(http.StatusBadRequest, w)
			return
		}

		//  -- Prep the data to send to the template -- //
		pageData := TopUsersData{
			TopList: top,
		}

		// Render template
		err := util.HttpServeTemplate("topusers.tmpl", pageData, w)
		if err != nil {
			log.Println("Error serving template:", err)
			util.HttpReturnError(http.StatusInternalServerError, w)
			return
		}

	} else if topType == "posts" {

		// Fetch top 10 liked posts
		top, err := post.GetTopPosts(10, basisCanvasId, since)
		if err != nil {
			log.Println("Error fetching top liked posts:", err)
			util.HttpReturnError(http.StatusInternalServerError, w)
			return
		}

		//  -- Prep the data to send to the template -- //
		pageData := TopPostsData{
			TopList: top,
		}

		// Render template
		err = util.HttpServeTemplate("topposts.tmpl", pageData, w)
		if err != nil {
			log.Println("Error serving template:", err)
			util.HttpReturnError(http.StatusInternalServerError, w)
			return
		}

	}

}
