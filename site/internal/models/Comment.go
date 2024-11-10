package models

import "time"

type Comment struct {
	UserId           int       `json:"user_id,omitempty"`
	PostId           int       `json:"post_id,omitempty"`
	Message          string    `json:"message,omitempty"`
	CreationDateTime time.Time `json:"creation_date_time,omitempty"`
}

type CommentExtended struct {
	Comment
	Commenter User `json:"commenter,omitempty"`
}
