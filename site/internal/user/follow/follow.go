package follow

import (
	"log"
	"rapidart/internal/database"
	"rapidart/internal/models"
)

// GetTopFollowedUsers fetches the top users by follower count
func GetTopFollowedUsers(limit int) ([]models.UserExtended, error) {
	users, err := database.GetUsersWithFollowerCountSortedByMostFollowers(limit)
	if err != nil {
		log.Println("Error fetching users with follower counts:", err)
		return nil, err
	}

	return users, nil
}
