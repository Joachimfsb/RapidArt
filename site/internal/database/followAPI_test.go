package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"rapidart/test"
	"testing"
)

func TestNewFollow(t *testing.T) {

	// Declare expectations
	//mock.ExpectCommit()
	user1 := test.GenTestUser()
	user2 := test.GenTestUser()

	follow := test.GenFollow(user1.UserId, user2.UserId)

	//mock.ExpectBegin()
	mock.ExpectExec(`^INSERT (.+)`).WithArgs(user1.UserId, user2.UserId).WillReturnResult(sqlmock.NewResult(1, 1))

	// Function call
	if err := NewFollow(follow); err != nil {
		t.Fatal("Got error trying to add follower: " + err.Error())
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}
