package user

import (
	"os"
	"rapidart/internal/database"
	"rapidart/internal/models"
	"rapidart/test"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

///////////////// TEST MAIN /////////////////

var mock sqlmock.Sqlmock // Internal mock used in database package for testing

// TestMain runs before every unit test in this package
func TestMain(m *testing.M) {

	mock = database.CreateMock()

	exitCode := m.Run()

	database.DeleteMock()

	os.Exit(exitCode)
}

////////////////// TESTS /////////////////////

func TestShouldCreateUser(t *testing.T) {

	u := test.GenTestUser()
	newUser := models.RegisterUser{
		Email:    u.Email,
		Username: u.Username,
		Password: "1234567890",
	}

	// Declare expectations
	mock.ExpectExec(`^INSERT (.+)`).WillReturnResult(sqlmock.NewResult(1, 1))

	// Function call
	if err := CreateUser(newUser); err != nil {
		t.Fatal("Got error trying to create user: " + err.Error())
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}

func TestShouldFailValidation(t *testing.T) {

	u := test.GenTestUser()
	newUser := models.RegisterUser{
		Email:    u.Email,
		Username: u.Username,
		Password: "tooshort",
	}

	// Declare expectations
	mock.ExpectExec(`^INSERT (.+)`).WillReturnResult(sqlmock.NewResult(1, 1))

	// Function call
	if err := CreateUser(newUser); err == nil {
		t.Fatal("CreateUser succeeded when it shouldn't!")
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err == nil {
		t.Fatal("Some expectations were met when they shouldn't")
	}
}
