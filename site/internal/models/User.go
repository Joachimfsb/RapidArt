package models

import "time"

type RapidUser struct {
	Username     string    `json:"username,omitempty"`
	Email        string    `json:"email,omitempty"`
	Displayname  string    `json:"displayname,omitempty"`
	Password     string    `json:"password,omitempty"`
	Passwordsalt string    `json:"passwordsalt,omitempty"`
	CreationTime time.Time `json:"creation_time"`
	Role         string    `json:"role,omitempty"`
	Bio          string    `json:"bio"`
	Profilepic   []byte    `json:"profilepic,omitempty"`
}
