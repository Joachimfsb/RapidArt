package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/database"
	"rapidart/internal/post"
	"rapidart/internal/post/like"
	"rapidart/internal/util"
	"strconv"
)

//!TEMPORARY DO NOT HAVE USER/client SAVE UserID, replace later with token/...

// Struct for saving post
type SavePostRequest struct {
	ImageData        string `json:"image_data"`
	BasisCanvasId    int    `json:"basis_canvas_id"`
	Caption          string `json:"caption"`
	TimeSpentDrawing int    `json:"time_spent_drawing"`
}

// POST request / saving post to database
func SavePost(w http.ResponseWriter, r *http.Request) {
	// Post req check
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// decode into SavePostRequest
	var req SavePostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// decode base64 into byte slice
	imageBytes, err := base64.StdEncoding.DecodeString(req.ImageData)
	if err != nil {
		fmt.Println("Error decoding base64 image:", err)
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Get session cookie
	cookie, err := r.Cookie("session-token")
	if err != nil {
		util.HttpReturnError(http.StatusUnauthorized, w)
		return
	}

	// Get currently logged in user
	session, err := auth.GetSession(cookie.Value)
	if err != nil {
		util.HttpReturnError(http.StatusUnauthorized, w)
		return
	}

	// save post to database
	id, err := post.CreatePost(session.UserId, req.BasisCanvasId, imageBytes, req.Caption, req.TimeSpentDrawing)
	if err != nil {
		fmt.Println("Error saving post to database:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(id))) // Return new id
}

// Gets image from database and serves it
func GetPost(w http.ResponseWriter, r *http.Request) {
	// Checks for post id
	if !r.URL.Query().Has("post_id") {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Make post id int
	postId, err := strconv.Atoi(r.URL.Query().Get("post_id"))
	if err != nil {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Get post from database by id
	post, err := database.GetPostById(postId)
	if err != nil {
		util.HttpReturnError(http.StatusNotFound, w)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Access-Control-Allow-Origin", "*") // All access for now, maybe change later
	w.Write(post.Image)                                // write image data to response, serving as BLOB
}

func PostLike(w http.ResponseWriter, r *http.Request) {

	// Parse params
	postIdStr := r.PathValue("id")
	if postIdStr == "" { // Missing post_id
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}
	postId, err := strconv.Atoi(postIdStr)
	if err != nil { // Bad format
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Get session
	session, err := auth.GetSession(util.GetSessionTokenFromCookie(r))
	if err != nil {
		util.HttpReturnError(http.StatusUnauthorized, w)
		return
	}

	// Add like to post
	err = like.LikePost(postId, session.UserId)
	if err != nil {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func PostUnlike(w http.ResponseWriter, r *http.Request) {

	// Parse params
	postIdStr := r.PathValue("id")
	if postIdStr == "" { // Missing post_id
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}
	postId, err := strconv.Atoi(postIdStr)
	if err != nil { // Bad format
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Get session
	session, err := auth.GetSession(util.GetSessionTokenFromCookie(r))
	if err != nil {
		util.HttpReturnError(http.StatusUnauthorized, w)
		return
	}

	// Add like to post
	success := like.UnlikePost(postId, session.UserId)
	if !success {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
