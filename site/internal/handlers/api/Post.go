package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"rapidart/internal/database"
	"rapidart/internal/post"
	"rapidart/internal/util"
	"strconv"
)

//!TEMPORARY DO NOT HAVE USER/client SAVE UserID, replace later with token/...

// Struct for saving post
type SavePostRequest struct {
	ImageData        string `json:"image_data"`
	BasisCanvasId    int    `json:"basis_canvas_id"`
	UserId           int    `json:"user_id"`
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

	// save post to database
	err = post.CreatePost(req.UserId, req.BasisCanvasId, imageBytes, req.Caption, req.TimeSpentDrawing)
	if err != nil {
		fmt.Println("Error saving post to database:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post saved"))
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
