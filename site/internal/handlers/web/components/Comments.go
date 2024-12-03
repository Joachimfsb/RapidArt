package components

import (
	"log"
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/models"
	"rapidart/internal/post/comment"
	"rapidart/internal/util"
	"sort"
	"strconv"
)

type commentsTemplateData struct {
	Comments     []models.CommentExtended
	LoggedInUser models.User
}

func Comments(w http.ResponseWriter, r *http.Request) {

	// Parse params
	postIdStr := r.PathValue("post_id")
	if postIdStr == "" { // Missing post_id
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}
	postId, err := strconv.Atoi(postIdStr)
	if err != nil { // Bad format
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Get logged in user
	loggedInUser, err := auth.GetLoggedInUser(util.GetSessionTokenFromCookie(r))
	if err != nil {
		util.HttpReturnError(http.StatusUnauthorized, w)
		return
	}

	// Get comments
	comments, err := comment.GetCommentsWithCommenterByPostId(postId)
	if err != nil { // error when fetching
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
	// Sort by recent first
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].CreationDateTime.After(comments[j].CreationDateTime)
	})

	commentsTemplateData := commentsTemplateData{
		Comments:     comments,
		LoggedInUser: loggedInUser,
	}

	if len(commentsTemplateData.Comments) > 0 {
		err = util.HttpServeTemplate("comments.tmpl", commentsTemplateData, w)
		if err != nil {
			log.Println(err)
			util.HttpReturnError(http.StatusInternalServerError, w)
			return
		}
	}
}
