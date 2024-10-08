package models

import "time"

// DB table BasisGallery
type BasisGallery struct {
	BasisGalleryId int
	StartDateTime  time.Time
	EndDateTime    time.Time
}
