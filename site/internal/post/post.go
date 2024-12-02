package post

import (
	"rapidart/internal/database"
	"rapidart/internal/models"
	"time"
)

// Create a new post
//
// Returns: post id (if created), error
func CreatePost(userId, basisCanvasId int, image []byte, caption string, timeSpent int) (int, error) {

	post := models.Post{
		UserId:           userId,
		BasisCanvasId:    basisCanvasId,
		Image:            image,
		Caption:          caption,
		TimeSpentDrawing: timeSpent,
		CreationDateTime: time.Now(),
		Active:           true,
	}

	id, err := database.AddPost(post)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func GetPostById(postId int) (models.Post, error) {
	return database.GetPostById(postId)
}

func GetRecentPostsByUser(userId int, limit uint) ([]models.PostExtended, error) {

	posts, err := database.GetPostsByUserId(userId, "creationDateTimeDesc", limit)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// Gets a users follows recent posts sorted by creation time (descending)
func GetRecentFollowsPosts(userId int, limit int) ([]models.PostExtended, error) {

	return database.GetUsersFollowsRecentPostsWithLikes(userId, limit, true)
}

// Gets recent posts sorted by creation time (descending)
func GetRecentPosts(limit int) ([]models.PostExtended, error) {

	return database.GetRecentPostsWithLikes(limit, true)
}
