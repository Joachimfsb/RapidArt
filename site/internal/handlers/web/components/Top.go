package components

import (
	"log"
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/models"
	"rapidart/internal/post"
	"rapidart/internal/post/comment"
	"rapidart/internal/post/like"
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

	// -- Get currently logged in user -- //
	// Get session cookie
	cookie, err := r.Cookie("session-token")
	if err != nil {
		util.HttpReturnError(http.StatusUnauthorized, w)
		return
	}

	// Get session
	session, err := auth.GetSession(cookie.Value)
	if err != nil {
		util.HttpReturnError(http.StatusUnauthorized, w)
		return
	}

	// --- Parse params --- //
	topType := r.PathValue("type")
	if topType == "" { // Missing post_id
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	if topType == "users" {
		///////////// TOP USERS //////////////

		// -- Parse query params -- //
		metric := r.URL.Query().Get("metric")
		if metric == "" {
			util.HttpReturnError(http.StatusBadRequest, w)
			return
		}

		// --- Fetch data --- //
		var top []models.UserExtended

		if metric == "likes" {

			// Fetch top 30 followed users
			var err error
			top, err = user.GetMostLikedUsers(30)
			if err != nil {
				log.Println("Error fetching top followed users:", err)
				util.HttpReturnError(http.StatusInternalServerError, w)
				return
			}

		} else if metric == "followers" {

			// Fetch top 30 followed users
			var err error
			top, err = follow.GetTopFollowedUsers(30)
			if err != nil {
				log.Println("Error fetching top followed users:", err)
				util.HttpReturnError(http.StatusInternalServerError, w)
				return
			}

		} else {
			util.HttpReturnError(http.StatusBadRequest, w)
			return
		}

		// Fill inn missing followercount and likecount
		for i, u := range top {
			stats, err := user.GetUserStats(u.UserId)
			if err != nil {
				util.HttpReturnError(http.StatusInternalServerError, w)
				return
			}

			top[i].TotalLikes = stats.TotalLikes
			top[i].FollowerCount = len(stats.Followers)
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
		////////////// TOP POSTS //////////////////

		// -- Parse query params -- //
		var since *time.Time = nil
		var basisCanvasId *int = nil

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

		// --- Fetch data --- //
		// Fetch top 30 liked posts
		top, err := post.GetTopPosts(30, basisCanvasId, since)
		if err != nil {
			log.Println("Error fetching top liked posts:", err)
			util.HttpReturnError(http.StatusInternalServerError, w)
			return
		}

		// Get their comments and check if user has liked
		for i, p := range top {
			// Comment count
			comments, err := comment.GetCommentsByPostId(p.PostId)
			if err != nil {
				util.HttpReturnError(http.StatusInternalServerError, w)
				return
			}

			top[i].CommentCount = len(comments)

			// Check if user has liked
			userHasLiked, err := like.HasUserLikedPost(session.UserId, p.PostId)
			if err != nil {
				util.HttpReturnError(http.StatusInternalServerError, w)
				return
			}

			top[i].UserHasLiked = userHasLiked
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

	} else {
		// Incorrect spesification of type
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

}
