package auth

import (
	crypto "rapidart/internal/crypto"
	"time"
)

type Session struct {
	UserId  int
	Expires time.Time
}

var sessions map[string]Session

func InitSessions() {
	sessions = make(map[string]Session)
}

func GetSession(token string) *Session {
	// Check if token exists in sessions
	val, ok := sessions[token]
	if ok {
		return &val // Yes? return the session
	} else {
		return nil // No? return nil
	}
}

func newSession(UserId int) string {
	// Generate token
	token := crypto.GenerateRandomCharacters(50)

	// If token exists already (very unlikely), try again recursively
	if GetSession(token) != nil {
		return newSession(UserId)
	} else {
		sessions[token] = Session{
			UserId:  UserId,
			Expires: time.Now().AddDate(0, 1, 0), // Expires 1 month after creation
		}
		return token
	}
}

func endSession(token string) {
	_, ok := sessions[token]
	if ok {
		delete(sessions, token)
	}
}
