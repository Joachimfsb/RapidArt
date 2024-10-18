package database

import (
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var mock sqlmock.Sqlmock

// a successful case
func InitMock() {
	var err error

	db, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	log.Println("Mock db initialized")
}

func CleanMock() {
	db.Close()

	log.Println("Mock db closed")
}

// TestMain sets up the firestore client to be used in each unit test.
func TestMain(m *testing.M) {

	InitMock()

	exitCode := m.Run()

	CleanMock()

	os.Exit(exitCode)
}
