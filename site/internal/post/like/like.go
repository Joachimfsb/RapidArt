package post

import (
	"log"
	"rapidart/internal/database"
	"rapidart/internal/models"
)

// Fetch top liked posts
func GetTopLikedPosts(limit int) ([]models.PostExtended, error) {
	posts, err := database.GetPostsWithLikeCountSortedByMostLikes(limit)
	if err != nil {
		log.Println("Error fetching posts with like counts:", err)
		return nil, err
	}

	return posts, nil
}

func GetUserProfilePic(userId int) ([]byte, error) {
	profilePic, err := database.GetUserProfilePic(userId)
	if err != nil {
		log.Println("Error fetching profile picture:", err)
		return nil, err
	}
	return profilePic, nil
}
