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

func GetRecentPostsByUser(userId int, limit uint) ([]models.PostExtended, error) {

	posts, err := database.GetPostsByUserId(userId, "creationDateTimeDesc", limit)
	if err != nil {
		return nil, err
	}

	return posts, nil
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
