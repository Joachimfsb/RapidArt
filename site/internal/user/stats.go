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
		FollowerCount: len(followers),
		FollowsCount:  len(follows),
		TotalLikes:    likes,
	}

	return stats, nil
}
