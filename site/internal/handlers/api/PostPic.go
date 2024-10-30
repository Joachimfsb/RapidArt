package api

import (
	"net/http"
	"rapidart/internal/post/postPic"
	"rapidart/internal/util"
	"strconv"
)

func ImgPostPic(w http.ResponseWriter, r *http.Request) {
	// Ensure that the method is GET
	if r.Method != http.MethodGet {
		util.HttpReturnError(http.StatusMethodNotAllowed, w)
		return
	}

	// Check if user_id is provided
	if !r.URL.Query().Has("post_id") {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Convert user_id
	postid, err := strconv.Atoi(r.URL.Query().Get("post_id"))
	if err != nil {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Fetch the profile picture
	postPic, err := postPic.GetPostPic(postid)
	if err != nil {
		util.HttpReturnError(http.StatusNotFound, w)
		return
	}

	// Set the content type as image/png
	w.Header().Set("Content-Type", "image/png")
	w.Write(postPic) // Serve the image
}
