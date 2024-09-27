package models

type RapidUser struct {
	UserId       int    `json:"user_id"`
	Username     string `json:"username,omitempty"`
	Email        string `json:"email,omitempty"`
	Displayname  string `json:"displayname,omitempty"`
	Password     string `json:"password,omitempty"`
	Passwordsalt string `json:"passwordsalt,omitempty"`
	CreationTime string `json:"creation_time"`
	Role         string `json:"role,omitempty"`
	Bio          string `json:"bio"`
	Profilepic   []byte `json:"profilepic,omitempty"`
}

type UserAuthentication struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
