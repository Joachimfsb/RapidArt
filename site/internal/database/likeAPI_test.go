package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"rapidart/internal/models"
	"testing"
)

func TestShouldAddLike(t *testing.T) {

	// Declare expectations
	//mock.ExpectBegin()
	mock.ExpectExec(`^INSERT (.+)`).WithArgs(3, 5).WillReturnResult(sqlmock.NewResult(1, 1))
	//mock.ExpectCommit()

	// Function call
	if err := AddLikeToPost(models.Like{
		UserId: 3,
		PostId: 5,
	}); err != nil {
		t.Fatal("Got error trying to add like: " + err.Error())
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}

func TestShouldGetTotalNumberOfLikesForEveryUserPost(t *testing.T) {
	// Test data
	uid := 1

	// Declare expectations
	mock.ExpectQuery("^SELECT COUNT").
		WithArgs(uid).
		WillReturnRows(sqlmock.NewRows([]string{"TotalLikes"}).AddRow(3))

	// Function call
	_, err := GetTotalLikesForEveryUserPost(uid)
	if err != nil {
		t.Fatal("Error returned from function call: " + err.Error())
	}
	// Check if expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}
