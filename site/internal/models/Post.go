package models

import "time"

// DB table Post
type Post struct {
	PostId           int       `json:"post_id"`
	UserId           int       `json:"user_id"`
	BasisCanvasId    int       `json:"basis_canvas_id"`
	Image            []byte    `json:"image"`
	Caption          string    `json:"caption"`
	TimeSpentDrawing int       `json:"time_spent_drawing"` // Milliseconds
	CreationDateTime time.Time `json:"creation_date_time"`
	Active           bool      `json:"active,omitempty"`
}

type PostExtended struct {
	Post
	LikeCount int `json:"like_count"`
}
