package models

import "time"

// DB table Session
type Session struct {
	SessionToken string    `json:"session_token"`
	UserId       int       `json:"user_id"`
	IPAddress    string    `json:"ip_address"`
	Browser      string    `json:"browser"`
	Expires      time.Time `json:"expires"`
}
