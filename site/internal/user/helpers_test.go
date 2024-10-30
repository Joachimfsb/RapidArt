package user

import (
	"rapidart/internal/models"
	"testing"
)

func TestShouldFailValidation(t *testing.T) {

	// Input data
	newUser := models.RegisterUser{
		Email:       "test@test.com",
		Username:    "test",
		Password:    "tooshort",
		Displayname: "Test",
	}

	// -- Function calls --

	// Too short password
	if validateRegistrationData(newUser) {
		t.Fatal("validateRegistrationData succeeded when it shouldn't: too short password")
	}

	// Email bad format
	newUser.Password = "1234567890"
	newUser.Email = "test"
	if validateRegistrationData(newUser) {
		t.Fatal("validateRegistrationData succeeded when it shouldn't: email invalid format")
	}

	// Username bad format
	newUser.Email = "test@test.com"
	newUser.Username = "ø"
	if validateRegistrationData(newUser) {
		t.Fatal("validateRegistrationData succeeded when it shouldn't: username invalid format")
	}

	// Display name bad / good format
	newUser.Password = "1234567890"
	newUser.Email = "test"
	newUser.Displayname = ""
	if validateRegistrationData(newUser) {
		t.Fatal("validateRegistrationData succeeded when it shouldn't: displayname invalid format")
	}
	newUser.Displayname = "12345678901234567890123456789012345678901234567890123456789012345678901"
	if validateRegistrationData(newUser) {
		t.Fatal("validateRegistrationData succeeded when it shouldn't: displayname invalid format")
	}
	newUser.Displayname = "1234567890123456789012345678901234567890123456789012345678901234567890"
	if validateRegistrationData(newUser) {
		t.Fatal("validateRegistrationData failed when it shouldn't: displayname good format")
	}
}
