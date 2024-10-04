package models

// DB table BasisCanvas
type BasisCanvas struct {
	BasisCanvasId  int
	BasisGalleryId int
	Type           string
	Image          []byte
}
