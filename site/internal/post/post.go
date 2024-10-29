package post

import (
	"errors"
	"log"
	"rapidart/internal/database"
	"rapidart/internal/models"
	"time"
)

// Create a new post
func CreatePost(userId, basisCanvasId int, image []byte, caption string, timeSpent int) error {

	post := models.Post{
		UserId:           userId,
		BasisCanvasId:    basisCanvasId,
		Image:            image,
		Caption:          caption,
		TimeSpentDrawing: timeSpent,
		CreationDateTime: time.Now(),
		Active:           true,
	}

	err := database.AddPost(post)
	if err != nil {
		return err
	}

	return nil
}

func CreateReport(report models.Report) error {
	const maxReports = 5
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
	if amountOfReports >= maxReports {
		err = database.DeactivateActivePost(report.PostId)
		if err != nil {
			return err
		}
	}
	return nil
}
