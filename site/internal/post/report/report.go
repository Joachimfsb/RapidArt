package report

import (
	"rapidart/internal/database"
	"rapidart/internal/models"
)

func AddReport(postId int, userId int, message string) error {
	report := models.Report{
		UserId:  userId,
		PostId:  postId,
		Message: message,
	}
	return database.NewReport(report)
}

func HasUserReportedPost(userId int, postId int) (bool, error) {
	return database.HasUserReportedPost(userId, postId)
}
