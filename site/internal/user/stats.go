package user

import (
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

	// Create model
	stats := models.UserStats{
		TotalLikes: likes,
	}

	return stats, nil
}
