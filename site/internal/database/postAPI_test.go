package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestShouldNotGetPostsByUserId(t *testing.T) {
	// Test data
	uid := 1
	var limit uint = 5

	// Declare expectations
	mock.ExpectQuery("SELECT").
		WithArgs(uid, limit).
		WillReturnRows(sqlmock.NewRows(nil))

	// Function call
	ret, err := GetPostsByUserId(uid, "creationDateTimeAsc", limit)
	if err != nil {
		t.Fatal("Error returned from function call!")
	}
	if ret != nil {
		t.Fatal("Got data when we shouldn't have!")
	}
	// Check if expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}
