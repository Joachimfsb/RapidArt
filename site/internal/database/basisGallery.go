package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"rapidart/internal/models"
	"time"
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

/**
 * It is important that when adding a new canvas, that you add a new basis gallery piece one at a time
 * so the new canvas gets the correct id
 */
func AddToGallery(newCanvas models.BasisGallery) error {
	//start time has to be set before this function
	newCanvas.EndDateTime = time.Now()

	sqlInsert := `
		INSERT INTO Basisgallery (
		                  StartTimestamp,
		                  EndTimestamp
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
