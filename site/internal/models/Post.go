package models

import "time"

// DB table Post
type Post2 struct {
	PostId           int
	UserId           int
	BasisCanvasId    int
	Image            []byte
	Caption          string
	TimeSpentDrawing int
	CreationDateTime time.Time
}

type Post struct {
	PostID            int       `json:"post_id,omitempty"`
	UserID            int       `json:"user_id,omitempty"`
	BasisCanvasID     int       `json:"basis_canvas_id,omitempty"`
	Image             []byte    `json:"image,omitempty"`
	Caption           string    `json:"caption"`
	TimeSpentDrawing  int       `json:"time_spent_drawing"`
	CreationTimestamp time.Time `json:"creation_timestamp,omitempty"`
	Active            bool      `json:"active,omitempty"`
}
