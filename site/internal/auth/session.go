package auth

import (
	"errors"
	"log"
	crypto "rapidart/internal/crypto"
	"rapidart/internal/database"
	"rapidart/internal/glob"
	"rapidart/internal/models"
	"rapidart/internal/util"
	"time"
)

///////////////// PUBLIC FUNCTIONS /////////////////////

// Get a session given a token
//
// Safe error messages
func GetSession(token string) (models.Session, error) {
	// Check if token exists in sessions
	session, err := database.GetSessionByToken(token)
	if err != nil {
		return models.Session{}, errors.New("session-not-found") // No? return error
	}

	// Yes? Check if it has not expired
	if time.Now().Before(session.Expires) {
		return session, nil // Not expired? return the session
	} else {
		// Has expired? Delete from array and return error
		endSession(token)

		return models.Session{}, errors.New("session-expired")
	}
}

////////////////////// PRIVATE FUNCTIONS ///////////////////////

// Create a new session for the user of the given id
func newSession(userId int, ipAddress string, browser string) string {
	// Generate token
	token := crypto.GenerateRandomCharacters(50)

	// If token exists already (very unlikely), try again recursively
	_, err := GetSession(token)
	if err == nil {
		return newSession(userId, ipAddress, browser)
	} else {
		// Add new session
		err := database.AddSession(models.Session{
			SessionToken: token,
			UserId:       userId,
			IPAddress:    ipAddress,
			Browser:      util.UserAgentToBrowser(browser),
			Expires:      time.Now().AddDate(0, 0, glob.SessionExpirationDays),
		})
		if err != nil {
			log.Println("Login failed for userId =", userId, " got error: ", err)
			return ""
		}
		return token
	}
}

// End a session for a given token
func endSession(token string) {
	database.DeleteSessionByToken(token)
}
