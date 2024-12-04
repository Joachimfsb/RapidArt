package models

// DB table like
type Like struct {
	UserId int `json:"user_id,omitempty"`
	PostId int `json:"post_id,omitempty"`
}
