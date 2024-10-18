package database

import (
	"fmt"
	"rapidart/internal/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestShouldGetUserById(t *testing.T) {

	var eu models.User // Expected user
	u := models.User{
		UserId:       1234, // Should not matter
		Username:     "testname",
		Email:        "test@test.com",
		Displayname:  "test namesson",
		Password:     "123abc",
		PasswordSalt: "123abc",
		CreationTime: time.Now(),
		Role:         "user",
		Bio:          "This is my bio",
		Profilepic:   nil,
	}

	rows := sqlmock.NewRows([]string{"UserId", "Username", "Email", "DisplayName", "PasswordHash", "PasswordSalt", "CreationDateTime", "Role", "Bio", "ProfilePicture"}).
		AddRow(u.UserId, u.Username, u.Email, u.Displayname, u.Password, u.PasswordSalt, u.CreationTime, u.Role, u.Bio, u.Profilepic)

	mock.ExpectQuery("^SELECT").WillReturnRows(rows)

	rs, _ := db.Query("SELECT")
	defer rs.Close()

	for rs.Next() {
		rs.Scan(eu.UserId, eu.Username, eu.Email, eu.Displayname, eu.Password, eu.PasswordSalt, eu.CreationTime, eu.Role, eu.Bio, eu.Profilepic)
	}

	if rs.Err() != nil {
		fmt.Println("got rows error:", rs.Err())
	}
}
