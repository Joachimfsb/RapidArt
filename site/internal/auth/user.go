package auth

import (
	"errors"
	"rapidart/internal/crypto"
	"rapidart/internal/database"
	"rapidart/internal/models"
)

// Perform a login for a given user.
// Either returns a token (token) or error statuses (wrongUser, wrongPass, err)
func Login(username string, password string) (token string, wrongUser bool, wrongPass bool, err error) {

	token = ""
	wrongUser = false
	wrongPass = false

	// Query user by username
	user, err := database.GetUserByUsername(username)
	if err != nil {
		wrongUser = true
		return // User not found
	}

	// Auth
	if user.Password != crypto.PBDKF2(password, user.PasswordSalt) {
		err = errors.New("password incorrect")
		wrongPass = true
		return // Password incorrect
	}

	// AUTH SUCCESS BELOW //

	// Generate session token
	token = newSession(user.UserId)
	return
}

// Perform logout for given token
func Logout(token string) {
	endSession(token)
}

// Gets the currently logged in user
// NOTE: If you only need the UserId, the function GetSession(...) should be preferred.
func GetLoggedInUser(token string) (models.User, error) {
	// Get session
	session, err := GetSession(token)
	if err != nil {
		return models.User{}, err
	}

	// Fetch user from DB
	user, err := database.GetUserById(session.UserId)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
