package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"rapidart/internal/models"
)

// Adds a new session (no checks)
func AddSession(session models.Session) error {

	sqlInsert := `
	INSERT INTO Session (
		SessionToken,
		UserId,
		IPAddress,
		Browser,
		Expires
	) VALUES (?, ?, ?, ?, ?);`

	_, err := db.Exec(sqlInsert,
		session.SessionToken,
		session.UserId,
		session.IPAddress,
		session.Browser,
		session.Expires,
	)
	if err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}

	return nil
}

// Gets a session given its token
func GetSessionByToken(token string) (models.Session, error) {

	var session models.Session

	row := db.QueryRow("SELECT SessionToken, UserId, IPAddress, Browser, Expires FROM Session WHERE SessionToken = ?", token)
	err := row.Scan(&session.SessionToken, &session.UserId, &session.IPAddress, &session.Browser, &session.Expires)

	if errors.Is(err, sql.ErrNoRows) {
		return models.Session{}, err
	}

	if err != nil {
		log.Println(err)
		return models.Session{}, err
	}

	// Convert times to local
	session.Expires = session.Expires.Local()

	return session, nil
}

func DeleteSessionByToken(token string) error {

	_, err := db.Exec("DELETE FROM Session WHERE SessionToken = ?;", token)
	if err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}

	return nil
}
