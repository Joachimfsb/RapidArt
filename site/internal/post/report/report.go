package report

import (
	"errors"
	"log"
	"rapidart/internal/database"
	"rapidart/internal/glob"
	"rapidart/internal/models"
	"time"
)

func AddReport(postId int, userId int, message string) error {
	report := models.Report{
		UserId:           userId,
		PostId:           postId,
		Message:          message,
		CreationDateTime: time.Now(),
	}

	// Save the report in the database
	err := database.NewReport(report)
	if err != nil {
		log.Println("NewReport error: [" + err.Error() + "]")
		return errors.New("server-error")
	}

	// Get the count of reports for the post
	amountOfReports, err := database.GetCountReports(postId)
	if err != nil {
		log.Println("Could not get count of reports for specified post id")
		return err
	}

	// Deactivate the post if the max reports (currently set to 5)
	if amountOfReports >= glob.NumberOfReportsBeforeDeactivatePost {
		err = database.DeactivateActivePost(postId)
		if err != nil {
			return err
		}
	}

	return nil
}

func HasUserReportedPost(userId int, postId int) (bool, error) {
	return database.HasUserReportedPost(userId, postId)
}
