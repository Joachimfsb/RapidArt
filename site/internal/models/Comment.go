package models

import "time"

// DB table Comment
type Comment struct {
	CommentId        int       `json:"comment_id"`
	UserId           int       `json:"user_id,omitempty"`
	PostId           int       `json:"post_id,omitempty"`
	Message          string    `json:"message,omitempty"`
	CreationDateTime time.Time `json:"creation_date_time,omitempty"`
}

// Extended comment information
type CommentExtended struct {
	Comment
	Commenter User `json:"commenter,omitempty"`
}
