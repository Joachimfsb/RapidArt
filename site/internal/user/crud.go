package user

import (
	"errors"
	"log"
	"rapidart/internal/database"
	"rapidart/internal/models"
)

// Validates input and creates a new user if vacant email and username
//
// Note: Error messages from this function are safe to display to the user.
func CreateUser(newUser models.RegisterUser) error {
	// Validate data
	if !validateRegistrationData(newUser) {
		return errors.New("validation-fail")
	}

	// Check email availability
	_, err := database.GetUserByEmail(newUser.Email)
	if err == nil {
		return errors.New("email-exists")
	}

	// Check username availability
	_, err = database.GetUserByUsername(newUser.Username)
	if err == nil {
		return errors.New("username-exists")
	}

	// Create new user
	err = database.AddUser(newUser.Email, newUser.Username, newUser.Password, "", "user", "", nil)
	if err != nil {
		log.Println("CreateUser error: [" + err.Error() + "]")
		return errors.New("server-error")
	}

	return nil
}

// TODO: Add functions CreateModerator, CreateAdmin
