package models

import "time"

type RapidUser struct {
	UserId       int       `json:"user_id"`
	Username     string    `json:"username,omitempty"`
	Email        string    `json:"email,omitempty"`
	Displayname  string    `json:"displayname,omitempty"`
	Password     string    `json:"password,omitempty"`
	PasswordSalt string    `json:"passwordsalt,omitempty"`
	CreationTime time.Time `json:"creation_time"`
	Role         string    `json:"role,omitempty"`
	Bio          string    `json:"bio"`
	Profilepic   []byte    `json:"profilepic,omitempty"`
}

type UserAuthentication struct {
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	PasswordSalt string `json:"password_salt,omitempty"`
}
