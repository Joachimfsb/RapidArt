package models

import "time"

// DB table BasisGallery
type BasisGallery struct {
	BasisGalleryId int
	StartTimestamp time.Time
	EndTimestamp   time.Time
}
