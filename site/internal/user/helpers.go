package user

import (
	"net/mail"
	"rapidart/internal/models"
	"regexp"
)

// Validates data in registration request.
func validateRegistrationData(u models.RegisterUser) bool {

	//////////////// EMAIL /////////////////

	if len(u.Email) == 0 || len(u.Email) > 255 {
		// Email empty or too long
		return false
	} else if _, err := mail.ParseAddress(u.Email); err != nil {
		// Bad email format
		return false
	}

	////////////// USERNAME ////////////////

	if len(u.Username) == 0 || len(u.Username) > 50 {
		// Empty or too long username
		return false
	} else if !regexp.MustCompile(`[a-zA-Z0-9]+`).MatchString(u.Username) {
		// Invalid format
		return false
	}

	////////////// PASSWORDS ///////////////

	if len(u.Password) < 10 || len(u.Password) > 255 {
		// Password is too short or too long
		return false
	}

	return true
}
