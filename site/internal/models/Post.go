package models

import "time"

// DB table Post
type Post struct {
	PostId           int       `json:"postId"`
	UserId           int       `json:"userId"`
	BasisCanvasId    int       `json:"basisCanvasId"`
	Image            []byte    `json:"image"`
	Caption          string    `json:"caption"`
	TimeSpentDrawing int       `json:"timeSpentDrawing"` // Milliseconds
	CreationDateTime time.Time `json:"creationDateTime"`
	Active           bool      `json:"active,omitempty"`
}

type PostExtended struct {
	Post
	LikeCount int `json:"like_count"`
}
