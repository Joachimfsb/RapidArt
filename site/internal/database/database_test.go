package database

import (
	"log"
	"os"
	"rapidart/test"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

/////////////// MOCK //////////////////

type Mock struct {
	M    sqlmock.Sqlmock
	Data MockData
}

type MockData struct {
	Users MockDataTuple
	Likes MockDataTuple
}

type MockDataTuple struct {
	Data interface{}
	Rows *sqlmock.Rows
}

// Create mock and setup fake db.
//
// Warning, only one mock can be in use at a time.
// Creation of multiple mock will result in only the last one working.
func CreateMock() *Mock {

	var err error
	var m sqlmock.Sqlmock

	db, m, err = sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	log.Println("Mock db initialized")

	mock := new(Mock)
	mock.M = m

	mock.populate()

	return mock
}

func (m *Mock) Delete() {
	db.Close()

	log.Println("Mock db closed")
}

// Populate the mock with data
func (m *Mock) populate() {

	//////// USER /////////
	u := test.GenTestUser()
	m.Data.Users.Data = u
	m.Data.Users.Rows = sqlmock.NewRows([]string{"UserId", "Username", "Email", "DisplayName", "PasswordHash", "PasswordSalt", "CreationDateTime", "Role", "Bio", "ProfilePicture"}).
		AddRow(u.UserId, u.Username, u.Email, u.Displayname, u.Password, u.PasswordSalt, u.CreationTime, u.Role, u.Bio, u.Profilepic)

}

///////////////// TEST /////////////////

var mock *Mock // Internal mock used in database package for testing

// TestMain runs before every unit test in this package
func TestMain(m *testing.M) {

	mock = CreateMock()

	exitCode := m.Run()

	mock.Delete()

	os.Exit(exitCode)
}
