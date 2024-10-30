package database

import (
	"database/sql/driver"
	"rapidart/test"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestShouldGetFollowersForUser(t *testing.T) {
	// Test data
	uid := 1

	// Declare expectations
	mock.ExpectQuery("^SELECT").
		WithArgs(uid).
		WillReturnRows(sqlmock.NewRows([]string{"f.FollowerUserId"}).AddRows(
			[]driver.Value{1},
			[]driver.Value{2},
			[]driver.Value{3},
		))

	// Function call
	followerList, err := GetFollowersForUser(uid)
	if err != nil {
		t.Fatal("Error returned from function call: " + err.Error())
	}
	if len(followerList) != 3 {
		t.Fatal("Incorrect number of followers returned: " + err.Error())
	}
	// Check if expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}

func TestShouldGetFollowsForUser(t *testing.T) {
	// Test data
	uid := 1

	// Declare expectations
	mock.ExpectQuery("^SELECT").
		WithArgs(uid).
		WillReturnRows(sqlmock.NewRows([]string{"f.FolloweeUserId"}).AddRows(
			[]driver.Value{1},
			[]driver.Value{2},
			[]driver.Value{3},
		))

	// Function call
	followsList, err := GetFollowsForUser(uid)
	if err != nil {
		t.Fatal("Error returned from function call: " + err.Error())
	}
	if len(followsList) != 3 {
		t.Fatal("Incorrect number of follows returned: " + err.Error())
	}
	// Check if expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}

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
