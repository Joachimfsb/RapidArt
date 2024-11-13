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
	"rapidart/internal/user"
	"rapidart/internal/util"
	"sort"
	"strconv"
)

type postTemplateData struct {
	LoggedInUser models.User
	BasisCanvas  models.BasisCanvas
	Post         models.Post
	Poster       models.User
	Comments     []models.CommentExtended
	LikeCount    int
	HasLiked     bool
	PosterIsSelf bool
}

func Post(w http.ResponseWriter, r *http.Request) {

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

	// Get post
	post, err := post.GetPostById(postId)
	if err != nil { // Post does not exist / error when fetching
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Get user who posted
	user, err := user.GetUserByUserId(post.UserId)
	if err != nil { // User does not exist (should not happen) / error when fetching (More likely)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Get comments
	comments, err := comment.GetCommentsWithCommenterByPostId(post.PostId)
	if err != nil { // error when fetching
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
	// Sort by recent first
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].CreationDateTime.After(comments[j].CreationDateTime)
	})

	// Get like count on post
	likeCount, err := like.GetNumberOfLikesOnPost(post.PostId)
	if err != nil { // error when fetching
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Get basis canvas used
	basisCanvas, err := basismanager.GetBasisCanvasById(post.BasisCanvasId)
	if err != nil { // error when fetching
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// User has liked the post
	hasLiked, err := like.HasUserLikedPost(loggedInUser.UserId, post.PostId)
	if err != nil { // error when fetching
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	posterIsSelf := false
	if post.UserId == loggedInUser.UserId {
		posterIsSelf = true
	}

	postTemplateData := postTemplateData{
		LoggedInUser: loggedInUser,
		BasisCanvas:  basisCanvas,
		Post:         post,
		Poster:       user,
		Comments:     comments,
		LikeCount:    likeCount,
		HasLiked:     hasLiked,
		PosterIsSelf: posterIsSelf,
	}

	err = util.HttpServeTemplate("post.tmpl", postTemplateData, w)
	if err != nil {
		log.Println(err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
