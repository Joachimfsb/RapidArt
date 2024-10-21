package models

import "time"

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

type UserExtended struct {
	User
	FollowerCount int `json:"follower_count"`
}

type RegisterUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
