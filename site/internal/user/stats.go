package user

import (
	"log"
	"rapidart/internal/database"
	"rapidart/internal/models"
)

func GetUserStats(userId int) (models.UserStats, error) {
	// Get like counts for all posts
	likes, err := database.GetTotalLikesForEveryUserPost(userId)
	if err != nil {
		return models.UserStats{}, err
	}

	// Get follows for user
	followers, err := database.GetFollowersForUser(userId)
	if err != nil {
		return models.UserStats{}, err
	}
	follows, err := database.GetFollowsForUser(userId)
	if err != nil {
		return models.UserStats{}, err
	}

	// Create model
	stats := models.UserStats{
		Followers:  followers,
		Follows:    follows,
		TotalLikes: likes,
	}

	return stats, nil
}

func GetMostLikedUsers(limit int) ([]models.UserExtended, error) {
	users, err := database.GetUsersWithMostTotalLikes(limit)
	if err != nil {
		log.Println("Error fetching users with follower counts:", err)
		return nil, err
	}

	return users, nil
}
