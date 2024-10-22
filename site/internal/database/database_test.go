package database

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

///////////////// TEST MAIN /////////////////

var mock sqlmock.Sqlmock // Internal mock used in database package for testing

// TestMain runs before every unit test in this package
func TestMain(m *testing.M) {

	mock = CreateMock()

	exitCode := m.Run()

	DeleteMock()

	os.Exit(exitCode)
}
