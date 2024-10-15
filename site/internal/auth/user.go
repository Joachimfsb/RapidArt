package auth

import (
	"errors"
	"rapidart/internal/crypto"
	"rapidart/internal/database"
)

func Login(username string, password string) (string, error) {
	// Query user by username
	user, err := database.GetUserByUsername(username)
	if err != nil {
		return "", err // User not found
	}

	// Auth
	if user.Password != crypto.PBDKF2(password, user.PasswordSalt) {
		return "", errors.New("password incorrect") // Password incorrect
	}

	// AUTH SUCCESS BELOW //

	// Generate session token
	return newSession(user.UserId), nil
}

func Logout(token string) {
	endSession(token)
}
