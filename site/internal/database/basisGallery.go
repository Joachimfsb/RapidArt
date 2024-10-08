package database

import (
	"database/sql"
	"errors"
	"fmt"
	"rapidart/internal/models"
)

// Fetches BasisGallery data based on the given ID.
//
// Note: if not found, an error is returned.
func GetBasisGalleryById(id int) (models.BasisGallery, error) {
	var gallery models.BasisGallery

	// Query DB
	row := db.QueryRow("SELECT BasisGalleryId, StartTimestamp, EndTimestamp FROM `BasisGallery` WHERE BasisGalleryId = ?", id)

	err := row.Scan(&gallery.BasisGalleryId, &gallery.StartTimestamp, &gallery.EndTimestamp)
	// No rows returned
	if errors.Is(err, sql.ErrNoRows) {
		return models.BasisGallery{}, fmt.Errorf("no basisgallery found by that id")
	}
	// General error
	if err != nil {
		return models.BasisGallery{}, err
	}

	// Convert times to local
	gallery.StartTimestamp = gallery.StartTimestamp.Local()
	gallery.EndTimestamp = gallery.EndTimestamp.Local()

	return gallery, nil
}
