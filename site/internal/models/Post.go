package models

import "time"

// DB table Post
type Post struct {
	PostID           int
	UserID           int
	BasisCanvasID    int
	Image            []byte
	Caption          string
	TimeSpentDrawing int
	CreationDateTime time.Time
	Active           bool `json:"active,omitempty"`
}
