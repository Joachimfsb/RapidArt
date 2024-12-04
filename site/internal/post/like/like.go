package like

import (
	"log"
	"rapidart/internal/database"
	"rapidart/internal/models"
)

// Return the number of likes on a post
//
// Returns: Number of likes, error
func GetNumberOfLikesOnPost(postId int) (int, error) {
	return database.GetLikeCountForPost(postId)
}

// Check if the user has liked a post
//
// Returns: Yes/No, error
func HasUserLikedPost(userId int, postId int) (bool, error) {
	return database.HasUserLikedPost(userId, postId)
}

// Like a post
//
// Returns: error
func LikePost(postId int, userId int) error {
	like := models.Like{
		UserId: userId,
		PostId: postId,
	}

	return database.AddLikeToPost(like)
}

// Unlike a post
//
// Returns: Success/Fail
func UnlikePost(postId int, userId int) bool {

	success, err := database.RemoveLikeFromPost(postId, userId)
	if err != nil {
		log.Println(err)
		return false
	}
	return success
}
