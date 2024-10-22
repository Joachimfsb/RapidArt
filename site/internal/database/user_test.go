package database

import (
	"rapidart/internal/models"
	"rapidart/test"
	"testing"
)

func TestShouldGetUserByUsername(t *testing.T) {

	// TEST DATA
	user := test.GenTestUser()

	mock.ExpectQuery("^SELECT").WithArgs(user.Username).WillReturnRows(GenRows([]models.User{user}))

	res, err := GetUserByUsername(user.Username)
	if err != nil {
		t.Fatal("Got error trying to get user by username: " + err.Error())
	}
	t.Log(res)
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}

func TestShouldGetUserById(t *testing.T) {

	user := test.GenTestUser()

	mock.ExpectQuery("^SELECT").WillReturnRows(GenRows([]models.User{user}))

	_, err := GetUserById(user.UserId)
	if err != nil {
		t.Fatal("Got error trying to get user by id: " + err.Error())
	}
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}
