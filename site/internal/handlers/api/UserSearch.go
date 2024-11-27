// In internal/handlers/api/search.go
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rapidart/internal/user"
	"rapidart/internal/util"
)

// user search api hanlder
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	// Get the query parameter
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing query parameter 'q'", http.StatusBadRequest)
		return
	}

	users, err := user.SearchUsers(query)
	if err != nil {
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Prepare the response data
	type UserResult struct {
		UserId        int    `json:"user_id"`
		Username      string `json:"username"`
		Displayname   string `json:"displayname"`
		ProfilePicURL string `json:"profile_pic_url"`
	}

	var results []UserResult
	for _, u := range users {
		picURL := fmt.Sprintf("/api/img/user/profile-pic/?userid=%d", u.UserId)
		results = append(results, UserResult{
			UserId:        u.UserId,
			Username:      u.Username,
			Displayname:   u.Displayname,
			ProfilePicURL: picURL,
		})
	}

	// Return the results as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(results)
	if err != nil {
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
