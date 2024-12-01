package web

import (
	"log"
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/basismanager"
	"rapidart/internal/models"
	"rapidart/internal/post"
	"rapidart/internal/post/comment"
	"rapidart/internal/post/like"
	"rapidart/internal/post/report"
	"rapidart/internal/user"
	"rapidart/internal/util"
	"time"
)

type postDetails struct {
	models.PostExtended
	Poster      models.User
	Comments    []models.CommentExtended
	HasLiked    bool
	HasReported bool
}

type IndexPageData struct {
	Title         string
	BasisCanvases []models.BasisCanvas
	Posts         []postDetails
}

func Index(w http.ResponseWriter, r *http.Request) {
	session, err := auth.GetSession(util.GetSessionTokenFromCookie(r))
	if err != nil {
		log.Println("Unauthorized access:", err)
		util.HttpReturnError(http.StatusUnauthorized, w)
		return
	}

	// Feed type (following or global)
	feedType := r.URL.Query().Get("feed")
	if feedType == "" {
		feedType = "followed"
	}

	// Get posts (30 of either recent or recent followed depending on the tab)
	var posts []models.PostExtended
	if feedType == "global" {
		posts, err = post.GetRecentPosts(30)
	} else {
		posts, err = post.GetRecentFollowsPosts(session.UserId, 30)
	}
	if err != nil {
		log.Println("Error fetching posts:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// More details for each post
	postsWithDetails := make([]postDetails, 0, len(posts))
	for _, p := range posts {
		// Get user who posted
		poster, err := user.GetUserByUserId(p.UserId)
		if err != nil {
			log.Println("Error fetching user for post:", err)
			continue
		}

		// Used to get number of comments
		comments, err := comment.GetCommentsWithCommenterByPostId(p.PostId)
		if err != nil {
			log.Println("Error fetching comments for post:", err)
			continue
		}

		// Like status
		hasLiked, err := like.HasUserLikedPost(session.UserId, p.PostId)
		if err != nil {
			log.Println("Error checking like status:", err)
			continue
		}

		// Report status
		hasReported, err := report.HasUserReportedPost(session.UserId, p.PostId)
		if err != nil {
			log.Println("Error checking report status:", err)
			continue
		}

		postsWithDetails = append(postsWithDetails, postDetails{
			PostExtended: p,
			Poster:       poster,
			Comments:     comments,
			HasLiked:     hasLiked,
			HasReported:  hasReported,
		})
	}

	// Basis canvases
	canvases, err := basismanager.GetBasisCanvasesByDateTime(time.Now())
	if err != nil {
		log.Println("Error fetching basis canvases:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Page data
	pageData := IndexPageData{
		Title:         "RapidArt - Front Page",
		BasisCanvases: canvases,
		Posts:         postsWithDetails,
	}

	// Render template
	err = util.HttpServeTemplate("index.tmpl", pageData, w)
	if err != nil {
		log.Println("Error serving template:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
	}
}
