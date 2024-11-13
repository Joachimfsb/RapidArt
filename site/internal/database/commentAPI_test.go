package database

import (
	"rapidart/test"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddCommentToPost(t *testing.T) {
	// Declare expectations
	//mock.ExpectCommit()
	user := test.GenTestUser()

	gallery := test.GenBasisGallery()
	canvas := test.GenBasisCanvas(gallery.BasisGalleryId)

	post, _ := test.GenTestPost(user.UserId, canvas.BasisCanvasId, false)
	comment := test.GenComment(user.UserId, post.PostId)

	//mock.ExpectBegin()
	mock.ExpectExec(`^INSERT (.+)`).WithArgs(comment.UserId, comment.PostId, comment.Message, comment.CreationDateTime).WillReturnResult(sqlmock.NewResult(1, 1))

	// Function call
	id, err := AddCommentToPost(comment)
	if err != nil {
		t.Fatal("Got error trying to add report: " + err.Error())
	}
	if id != 1 {
		t.Fatal("Wrong id returned")
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}

// inspirert av TestGetAllReportsForPost(...)
func TestGetAllCommentsFromPost(t *testing.T) {
	// Generate test data
	user := test.GenTestUser()
	gallery := test.GenBasisGallery()
	canvas := test.GenBasisCanvas(gallery.BasisGalleryId)
	post, _ := test.GenTestPost(user.UserId, canvas.BasisCanvasId, false)
	comment := test.GenComment(user.UserId, post.PostId)

	// Now expect the SELECT query for GetAllReportsForPost
	rows := sqlmock.NewRows([]string{"CommentId", "UserId", "PostId", "Message", "CreationDateTime"}).
		AddRow(comment.CommentId, comment.UserId, comment.PostId, comment.Message, comment.CreationDateTime)

	mock.ExpectQuery(`^SELECT \* FROM Comment WHERE PostId = ?`).WithArgs(comment.PostId).WillReturnRows(rows)

	// Call GetAllReportsForPost to execute the SELECT expectation
	comments, err := GetAllCommentsFromPost(comment.PostId)
	if err != nil {
		t.Fatal("Got error trying to count comments: " + err.Error())
	}

	// Check that the returned report amount is as expected
	if len(comments) != 1 {
		t.Fatal("Wrong amount of reports: expected 1, got", len(comments))
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}
