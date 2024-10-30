package user

import (
	"rapidart/internal/models"
	"regexp"
)

// Validates data in registration request.
func validateRegistrationData(u models.RegisterUser) bool {

	//////////////// EMAIL /////////////////

	if len(u.Email) == 0 || len(u.Email) > 255 {
		// Email empty or too long
		return false
	} else if !regexp.MustCompile(`^(?:[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])?`).MatchString(u.Email) {
		// Bad email format
		return false
	}

	////////////// USERNAME ////////////////

	if len(u.Username) == 0 || len(u.Username) > 50 {
		// Empty or too long username
		return false
	} else if !regexp.MustCompile(`^[a-z0-9]+$`).MatchString(u.Username) {
		// Invalid format
		return false
	}

	////////////// PASSWORDS ///////////////

	if len(u.Password) < 10 || len(u.Password) > 255 {
		// Password is too short or too long
		return false
	}

	////////////// DISPLAY NAME //////////////

	if len(u.Displayname) == 0 || len(u.Displayname) > 70 {
		// Display name is too short or too long
		return false
	}

	return true
}
