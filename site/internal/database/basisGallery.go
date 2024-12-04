package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"rapidart/internal/models"
)

// Fetches BasisGallery data based on the given ID.
//
// Note: if not found, an error is returned.
func GetBasisGalleryById(id int) (models.BasisGallery, error) {
	var gallery models.BasisGallery

	// Query DB
	row := db.QueryRow("SELECT BasisGalleryId, StartDateTime, EndDateTime FROM `BasisGallery` WHERE BasisGalleryId = ?", id)

	err := row.Scan(&gallery.BasisGalleryId, &gallery.StartDateTime, &gallery.EndDateTime)
	// No rows returned
	if errors.Is(err, sql.ErrNoRows) {
		return models.BasisGallery{}, fmt.Errorf("no basisgallery found by that id")
	}
	// General error
	if err != nil {
		return models.BasisGallery{}, err
	}

	// Convert times to local
	gallery.StartDateTime = gallery.StartDateTime.Local()
	gallery.EndDateTime = gallery.EndDateTime.Local()

	return gallery, nil
}

// Adds a new basisGallery
func AddGallery(newCanvas models.BasisGallery) error {

	sqlInsert := `
		INSERT INTO Basisgallery (
		                  StartDateTime,
		                  EndDateTime
		) VALUES (?, ?);`

	_, err := db.Exec(sqlInsert,
		newCanvas.StartDateTime,
		newCanvas.EndDateTime,
	)
	if err != nil {
		log.Println("Error: ", err)
		fmt.Println(err)
		return fmt.Errorf("ERROR: %v", err)
	}

	return nil
}
