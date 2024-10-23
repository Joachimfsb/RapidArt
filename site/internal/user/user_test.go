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

///////////////// TEST HELPERS /////////////////

////////////////// TESTS /////////////////////

func TestShouldCreateUser(t *testing.T) {

	u := test.GenTestUser()
	newUser := models.RegisterUser{
		Email:    u.Email,
		Username: u.Username,
		Password: "1234567890",
	}

	// Declare db expectations
	mock.ExpectQuery(`^SELECT`).WithArgs(newUser.Email).WillReturnRows(sqlmock.NewRows(nil))
	mock.ExpectQuery(`^SELECT`).WithArgs(newUser.Username).WillReturnRows(sqlmock.NewRows(nil))
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

func TestShouldFailChecks(t *testing.T) {

	// Input data
	u := test.GenTestUser()
	u.Email = "test@test.com"
	u.Username = "test"
	newUser := models.RegisterUser{
		Email:    u.Email,
		Username: u.Username + "notamatch",
		Password: "tooshort",
	}

	rows := database.GenRows([]models.User{u}) // Describe rows that the db should return

	//////////////////////////////////////////////////
	t.Log("TEST CASE 1: Email exists")

	// -- Declare expectations --
	// Explained: If db gets a query starting with "SELECT", with the arguments newUser.Email, return the given rows
	mock.ExpectQuery(`^SELECT`).WithArgs(newUser.Email).WillReturnRows(rows) // Match

	// -- Function calls --
	// Test the function
	if err := CreateUser(newUser); err == nil {
		t.Fatal("CreateUser succeeded when it shouldn't")
	}

	// -- we make sure that all expectations were met --
	if err := mock.ExpectationsWereMet(); err == nil {
		t.Fatal("Some mock expectations were met when they shouldn't")
	}

	////////////////////////////////////////////////
	t.Log("TEST CASE 1: Username exists")

	newUser.Email = u.Email + "notamatch"
	newUser.Username = u.Username

	// -- Declare expectations --
	mock.ExpectQuery(`^SELECT`).WithArgs(newUser.Email).WillReturnRows(sqlmock.NewRows(nil))
	mock.ExpectQuery(`^SELECT`).WithArgs(newUser.Username).WillReturnRows(rows) // Match

	// -- Function calls --
	if err := CreateUser(newUser); err == nil {
		t.Fatal("CreateUser succeeded when it shouldn't")
	}

	// -- we make sure that all expectations were met --
	if err := mock.ExpectationsWereMet(); err == nil {
		t.Fatal("Some mock expectations were met when they shouldn't")
	}
}
