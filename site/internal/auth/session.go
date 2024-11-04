package auth

import (
	"errors"
	crypto "rapidart/internal/crypto"
	"time"
)

type Session struct {
	UserId  int
	Expires time.Time
}

var sessions map[string]Session

///////////////// PUBLIC FUNCTIONS /////////////////////

// Run this at startup to initialize data structure
func InitSessions() {
	sessions = make(map[string]Session)
}

// Get a session given a token
//
// Safe error messages
func GetSession(token string) (Session, error) {
	// Check if token exists in sessions
	val, ok := sessions[token]
	if ok {
		// Yes? Check if it has not expired
		if time.Now().Before(val.Expires) {
			return val, nil // Not expired? return the session
		} else {
			// Has expired? Delete from array and return error
			endSession(token)

			return Session{}, errors.New("session-expired")
		}
	} else {
		return Session{}, errors.New("session-not-found") // No? return error
	}
}

////////////////////// PRIVATE FUNCTIONS ///////////////////////

// Create a new session for the user of the given id
func newSession(UserId int) string {
	// Generate token
	token := crypto.GenerateRandomCharacters(50)

	// If token exists already (very unlikely), try again recursively
	_, err := GetSession(token)
	if err == nil {
		return newSession(UserId)
	} else {
		// Add new session
		sessions[token] = Session{
			UserId:  UserId,
			Expires: time.Now().AddDate(0, 3, 0), // Expires 3 month after creation
		}
		return token
	}
}

// End a session for a given token
func endSession(token string) {
	_, ok := sessions[token]
	if ok {
		delete(sessions, token)
	}
}
