package database

import (
	"rapidart/test"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestNewReport(t *testing.T) {

	// Declare expectations
	//mock.ExpectCommit()
	user := test.GenTestUser()

	gallery := test.GenBasisGallery()
	canvas := test.GenBasisCanvas(gallery.BasisGalleryId)

	post, _ := test.GenTestPost(user.UserId, canvas.BasisCanvasId, false)
	report := test.GenReport(user.UserId, post.PostId)

	mock.ExpectExec(`^INSERT (.+)`).WithArgs(report.UserId, report.PostId, report.Message, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1)) //this line is from ChatGPT
	// Function call
	if err := NewReport(report); err != nil {
		t.Fatal("Got error trying to add report: " + err.Error())
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}

// Laget av/ved hjelp av KI/AI
// https://chatgpt.com/share/671fa6cb-7a1c-800c-a376-6f0851377a9c
func TestGetAllReportsForPost(t *testing.T) {
	// Generate test data
	user1 := test.GenTestUser()
	user2 := test.GenTestUser()
	gallery := test.GenBasisGallery()
	canvas := test.GenBasisCanvas(gallery.BasisGalleryId)
	post, _ := test.GenTestPost(user1.UserId, canvas.BasisCanvasId, false)
	report := test.GenReport(user2.UserId, post.PostId)

	// Now expect the SELECT query for GetAllReportsForPost
	rows := sqlmock.NewRows([]string{"UserId", "PostId", "Message", "CreationDateTime"}).
		AddRow(report.UserId, report.PostId, report.Message, report.CreationDateTime)

	mock.ExpectQuery(`^SELECT \* FROM Report WHERE PostId = ?`).WithArgs(report.PostId).WillReturnRows(rows)

	// Call GetAllReportsForPost to execute the SELECT expectation
	reportAmount, err := GetAllReportsForPost(report.PostId)
	if err != nil {
		t.Fatal("Got error trying to count reports: " + err.Error())
	}

	// Check that the returned report amount is as expected
	if len(reportAmount) != 1 {
		t.Fatal("Wrong amount of reports: expected 1, got", len(reportAmount))
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}
