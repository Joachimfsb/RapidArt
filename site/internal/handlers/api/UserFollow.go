package api

import (
	"net/http"
	"rapidart/internal/auth"
	"rapidart/internal/user/follow"
	"rapidart/internal/util"
	"strconv"
)

// Adds or removes a follow to another user
func UserFollow(w http.ResponseWriter, r *http.Request) {

	// Parse params
	// User id
	userIdStr := r.PathValue("UserId")
	if userIdStr == "" { // Missing user id
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil { // Bad format
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Value
	valueStr := r.PathValue("Value")
	if valueStr == "" { // Missing user id
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}
	value, err := strconv.Atoi(valueStr)
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

	// Make sure user isn't trying to follow themselves
	if session.UserId == userId {
		util.HttpReturnError(http.StatusForbidden, w)
		return
	}

	var res bool
	switch value {
	case 1:
		// Add follow to user
		res = follow.Follow(session.UserId, userId)
	case 0:
		// Add follow to user
		res = follow.UnFollow(session.UserId, userId)
	default:
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Could not add / remove follow
	if !res {
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
