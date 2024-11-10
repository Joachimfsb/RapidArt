package like

import (
	"log"
	"rapidart/internal/database"
	"rapidart/internal/models"
)

func GetNumberOfLikesOnPost(postId int) (int, error) {
	return database.GetLikeCountForPost(postId)
}

func HasUserLikedPost(userId int, postId int) (bool, error) {
	return database.HasUserLikedPost(userId, postId)
}

// Fetch top liked posts
func GetTopLikedPosts(limit int) ([]models.PostExtended, error) {
	posts, err := database.GetPostsWithLikeCountSortedByMostLikes(limit)
	if err != nil {
		log.Println("Error fetching posts with like counts:", err)
		return nil, err
	}

	return posts, nil
}
