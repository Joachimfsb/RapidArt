package auth

import (
	"errors"
	"rapidart/internal/crypto"
	"rapidart/internal/database"
)

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

func Logout(token string) {
	endSession(token)
}
