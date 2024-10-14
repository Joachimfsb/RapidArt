package models

import "time"

// DB table Post
type Post struct {
	PostId           int
	UserId           int
	BasisCanvasId    int
	Image            []byte
	Caption          string
	TimeSpentDrawing int
	CreationDateTime time.Time
}
