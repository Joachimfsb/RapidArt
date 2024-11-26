package follow

import (
	"log"
	"rapidart/internal/database"
	"rapidart/internal/models"
)

// GetTopFollowedUsers fetches the top users by follower count
//
// The following fields are populated: UserId, Username, DisplayName, ProfilePicture, FollowerCount
func GetTopFollowedUsers(limit int) ([]models.UserExtended, error) {
	users, err := database.GetUsersWithFollowerCountSortedByMostFollowers(limit)
	if err != nil {
		log.Println("Error fetching users with follower counts:", err)
		return nil, err
	}

	return users, nil
}

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

func UnFollow(followerId int, followeeId int) bool {
	success, err := database.RemoveFollow(followerId, followeeId)
	if err != nil {
		log.Println(err)
		return false
	}
	return success
}
