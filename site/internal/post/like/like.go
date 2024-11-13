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

func LikePost(postId int, userId int) error {
	like := models.Like{
		UserId: userId,
		PostId: postId,
	}

	return database.AddLikeToPost(like)
}

func UnlikePost(postId int, userId int) bool {

	success, err := database.RemoveLikeFromPost(postId, userId)
	if err != nil {
		log.Println(err)
		return false
	}
	return success
}
