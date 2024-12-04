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

// Get a post by its id.
//
// Returns: post, error
func GetPostById(postId int) (models.Post, error) {
	return database.GetPostById(postId)
}

// Get a users recent posts.
//
// Returns: Post with LikeCount, error
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

// Get top posts gets the top liked posts given certain optional filters.
//
// limit: limits the number of results.
//
// basisCanvasFilter: Filter the results on a given basiscanvas
//
// sinceFilter: Show only results since the time specified in this variable
//
// Returns: list of posts with post data and likecount, error
func GetTopPosts(limit int, basisCanvasFilter *int, sinceFilter *time.Time) ([]models.PostExtended, error) {
	if basisCanvasFilter != nil {
		// Prep time
		newSinceFilter := time.Date(2000, 0, 0, 0, 0, 0, 0, time.Local) // The dawn of time
		if sinceFilter != nil {
			newSinceFilter = *sinceFilter
		}

		return database.GetPostsWithLikeCountSortedByMostLikesFilteredByBasisCanvasIdAndSince(limit, *basisCanvasFilter, newSinceFilter)
	} else {
		if sinceFilter != nil {
			// Only since filter
			return database.GetPostsWithLikeCountSortedByMostLikesFilteredBySince(limit, *sinceFilter)
		} else {
			// No filters specified
			return database.GetPostsWithLikeCountSortedByMostLikes(limit)
		}
	}
}
