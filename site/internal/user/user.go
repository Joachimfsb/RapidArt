package user

import (
	"errors"
	"log"
	"rapidart/internal/crypto"
	"rapidart/internal/database"
	"rapidart/internal/models"
	"strings"
	"time"
)

// Validates input and creates a new user if vacant email and username
//
// Note: Error messages from this function are safe to display to the user.
func CreateUser(newUser models.RegisterUser) error {

	// Convert email and username to lower case
	newUser.Email = strings.ToLower(newUser.Email)
	newUser.Username = strings.ToLower(newUser.Username)

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
	salt := crypto.GenerateRandomCharacters(5)

	userModel := models.User{
		Email:        newUser.Email,
		Username:     newUser.Username,
		Displayname:  newUser.Displayname,
		PasswordSalt: salt,
		Password:     crypto.PBDKF2(newUser.Password, salt),
		CreationTime: time.Now(),
		Role:         "user",
		Bio:          "",
		Profilepic:   newUser.ProfilePic,
	}

	// Add to db
	err = database.AddUser(userModel)
	if err != nil {
		log.Println("CreateUser error: [" + err.Error() + "]")
		return errors.New("server-error")
	}

	return nil
}

// Fetches the user info of the given user
func GetUserByUsername(username string) (models.User, error) {
	return database.GetUserByUsername(username)
}

// Checks whether an email is available to be registered.
//
// Returns: available = true, taken = false
func CheckEmailAvailability(email string) bool {
	_, err := database.GetUserByEmail(email)
	return err != nil
}

// Checks whether a username is available to be registered.
//
// Returns: available = true, taken = false
func CheckUsernameAvailability(username string) bool {
	_, err := database.GetUserByUsername(username)
	return err != nil
}
