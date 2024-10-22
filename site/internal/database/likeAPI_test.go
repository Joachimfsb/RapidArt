package database

import (
	"rapidart/internal/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// a successful case
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
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
