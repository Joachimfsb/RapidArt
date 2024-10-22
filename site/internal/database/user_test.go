package database

import (
	"rapidart/internal/models"
	"testing"
)

func TestShouldGetUserById(t *testing.T) {

	u := mock.Data.Users.Data.(models.User)
	mock.M.ExpectQuery("^SELECT").WillReturnRows(mock.Data.Users.Rows)

	_, err := GetUserById(u.UserId)
	if err != nil {
		t.Fatal("Got error trying to get user by id: " + err.Error())
	}
	err = mock.M.ExpectationsWereMet()
	if err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}
