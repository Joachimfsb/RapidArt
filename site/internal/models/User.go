package models

import (
	"time"
)

// DB table User
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

// ///////// HELPER STRUCTS ////////////
// Extended user information
type UserExtended struct {
	User
	FollowerCount int `json:"follower_count"`
	TotalLikes    int `json:"like_count"`
}

// User statistics
type UserStats struct {
	Followers  []int `json:"followers"`
	Follows    []int `json:"follows"`
	TotalLikes int   `json:"total_likes"`
}

// User registration
type RegisterUser struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Displayname string `json:"displayname"`
	ProfilePic  []byte `json:"profile_pic"`
}
