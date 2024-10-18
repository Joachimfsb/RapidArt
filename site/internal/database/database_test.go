package database

import (
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

/////////////// PUBLIC //////////////////

// a successful case
func CreateMock() sqlmock.Sqlmock {
	var err error
	var mock sqlmock.Sqlmock

	db, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	log.Println("Mock db initialized")
	return mock
}

func DeleteMock() {
	db.Close()

	log.Println("Mock db closed")
}

///////////////// PRIVATE /////////////////

var mock sqlmock.Sqlmock

// TestMain runs before every unit test in this package
func TestMain(m *testing.M) {

	mock = CreateMock()

	exitCode := m.Run()

	os.Exit(exitCode)
}
