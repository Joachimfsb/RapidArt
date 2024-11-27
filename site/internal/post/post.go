package post

import (
	"errors"
	"log"
	"rapidart/internal/database"
	"rapidart/internal/glob"
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

func CreateReport(report models.Report) error {

	report = models.Report{
		UserId:           report.UserId,
		PostId:           report.PostId,
		Message:          report.Message,
		CreationDateTime: time.Now(),
	}
	err := database.NewReport(report)
	if err != nil {
		log.Println("NewReport error: [" + err.Error() + "]")
		return errors.New("server-error")
	}

	amountOfReports, err := database.GetCountReports(report.PostId)
	if err != nil {
		log.Println("Could not get count of reports for specified post id")
		return err
	}
	if amountOfReports >= glob.MaxReports {
		err = database.DeactivateActivePost(report.PostId)
		if err != nil {
			return err
		}
	}
	return nil
}
