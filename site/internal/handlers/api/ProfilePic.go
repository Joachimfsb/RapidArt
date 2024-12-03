package api

import (
	"net/http"
	"rapidart/internal/user/profile"
	"rapidart/internal/util"
	"strconv"
)

func ImgUserProfilePic(w http.ResponseWriter, r *http.Request) {

	// Check if user_id is provided
	if !r.URL.Query().Has("userid") {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Convert user_id
	userId, err := strconv.Atoi(r.URL.Query().Get("userid"))
	if err != nil {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Fetch the profile picture
	profilePicture, err := profile.GetUserProfilePic(userId)
	if err != nil {
		util.HttpReturnError(http.StatusNotFound, w)
		return
	}

	// Set the content type as image/png
	w.Header().Set("Content-Type", "image/png")
	w.Write(profilePicture) // Serve the image
}
