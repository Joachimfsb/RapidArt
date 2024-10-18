package models

import (
	"bytes"
	"time"
)

type User struct {
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

func (u User) Equals(other User) bool {
	return (u.UserId == other.UserId &&
		u.Username == other.Username &&
		u.Email == other.Email &&
		u.Displayname == other.Displayname &&
		u.Password == other.Password &&
		u.PasswordSalt == other.PasswordSalt &&
		u.CreationTime == other.CreationTime &&
		u.Bio == other.Bio &&
		bytes.Equal(u.Profilepic, other.Profilepic))
}

/////////// HELPER STRUCTS ////////////

type RegisterUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
