package follow

import (
	"log"
	"rapidart/internal/database"
	"rapidart/internal/models"
)

// Add a new follow
//
// Returns: Success/Fail
func Follow(followerId int, followeeId int) bool {
	follow := models.Follow{
		FollowerUserId: followerId,
		FolloweeUserId: followeeId,
	}

	err := database.NewFollow(follow)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// Remove a follow
//
// Returns: Success/Fail
func UnFollow(followerId int, followeeId int) bool {
	success, err := database.RemoveFollow(followerId, followeeId)
	if err != nil {
		log.Println(err)
		return false
	}
	return success
}
