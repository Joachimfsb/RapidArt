package follow

import (
	"log"
	"rapidart/internal/database"
	"rapidart/internal/models"
)

// GetTopFollowedUsers fetches the top users by follower count
func GetTopFollowedUsers() ([]models.User, error) {
	rows, err := database.GetUsersWithFollowerCounts()
	if err != nil {
		log.Println("Error fetching users with follower counts:", err)
		return nil, err
	}
	defer rows.Close()

	// Process rows into User structs
	var topUsers []models.User
	for rows.Next() {
		var user models.User
		var followerCount int
		// Scan the row
		err := rows.Scan(&user.UserId, &user.Displayname, &user.Profilepic, &followerCount)
		if err != nil {
			log.Println("Error scanning user data:", err)
			continue
		}
		user.FollowerCount = followerCount
		topUsers = append(topUsers, user)
	}

	// Return the processed list of top users
	return topUsers, nil
}
