package database

import (
	"fmt"
	"rapidart/internal/models"
	"rapidart/test"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestShouldGetUserById(t *testing.T) {

	u := test.GenTestUser() // Expected user
	var ret models.User

	rows := sqlmock.NewRows([]string{"UserId", "Username", "Email", "DisplayName", "PasswordHash", "PasswordSalt", "CreationDateTime", "Role", "Bio", "ProfilePicture"}).
		AddRow(u.UserId, u.Username, u.Email, u.Displayname, u.Password, u.PasswordSalt, u.CreationTime, u.Role, u.Bio, u.Profilepic)

	mock.ExpectQuery("^SELECT").WillReturnRows(rows)

	rs, _ := db.Query("SELECT")
	defer rs.Close()

	for rs.Next() {
		rs.Scan(ret.UserId, ret.Username, ret.Email, ret.Displayname, ret.Password, ret.PasswordSalt, ret.CreationTime, ret.Role, ret.Bio, ret.Profilepic)
	}

	if u.Equals(ret) {
		t.Fatal("User didn't match expected user")
	}

	if rs.Err() != nil {
		fmt.Println("got rows error:", rs.Err())
	}
}
