package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"rapidart/internal/models"
	"time"
)

// Fetches a list of BasisCanvases given a date and time. Fetches this via the BasisGallery table.
//
// Note: if none found, NO ERROR IS RETURNED - Only an empty list
func GetBasisCanvasesByDateTime(datetime time.Time) ([]models.BasisCanvas, error) {
	var canvases []models.BasisCanvas

	// Query DB
	rows, err := db.Query(
		"SELECT BC.BasisCanvasId, BC.BasisGalleryId, BC.Type, BC.Image "+
			"FROM `BasisGallery` as BG "+
			"INNER JOIN `BasisCanvas` as BC ON BG.BasisGalleryId = BC.BasisGalleryId "+
			"WHERE ? BETWEEN BG.StartDateTime AND BG.EndDateTime",
		datetime)
	// General error
	if err != nil {
		return []models.BasisCanvas{}, err
	}

	defer rows.Close()

	// Loop through rows
	for rows.Next() {
		canvases = append(canvases, models.BasisCanvas{})

		err := rows.Scan(
			&canvases[len(canvases)-1].BasisCanvasId,
			&canvases[len(canvases)-1].BasisGalleryId,
			&canvases[len(canvases)-1].Type,
			&canvases[len(canvases)-1].Image,
		)
		// General error
		if err != nil {
			return []models.BasisCanvas{}, err
		}
	}

	return canvases, nil
}

// Fetches BasisCanvas data based on the given ID.
//
// If not found, an error is returned.
func GetBasisCanvasById(id int) (models.BasisCanvas, error) {
	var canvas models.BasisCanvas

	// Query DB
	row := db.QueryRow("SELECT BasisCanvasId, BasisGalleryId, Type, Image FROM `BasisCanvas` WHERE BasisCanvasId = ?", id)

	err := row.Scan(&canvas.BasisCanvasId, &canvas.BasisGalleryId, &canvas.Type, &canvas.Image)
	// No rows returned
	if errors.Is(err, sql.ErrNoRows) {
		return models.BasisCanvas{}, fmt.Errorf("no basiscanvas found by that id")
	}
	// General error
	if err != nil {
		return models.BasisCanvas{}, err
	}

	return canvas, nil
}

/**
 * It is important that when adding a new canvas, that you add a new basis gallery piece one at a time
 * so the new canvas gets the correct id
 */
func AddNewCanvas(newBasisCanvas models.BasisCanvas) error {

	var galleryIdExist int
	err := db.QueryRow("SELECT COUNT(1) FROM Basisgallery WHERE BasisGalleryId=?", newBasisCanvas.BasisGalleryId).Scan(&galleryIdExist)
	if err != nil {
		log.Println("Error checking gallery ID:", err)
		return err
	}

	if galleryIdExist == 0 {
		log.Println("No gallery with this id")
		return fmt.Errorf("gallery ID does not exist")
	}

	/*var count = 0
	count, err := HowManyGallery()
	if err != nil {
		log.Println(err)
		return err
	}
	newBasisCanvas.BasisGalleryId = count*/
	sqlInsert := `
		INSERT INTO Basiscanvas (
		                  BasisGalleryId,
		                  Type,
		                  Image
		) VALUES (?, ?, ?);`

	_, err = db.Exec(sqlInsert,
		newBasisCanvas.BasisGalleryId,
		newBasisCanvas.Type,
		newBasisCanvas.Image,
	)
	if err != nil {
		log.Println("Error: ", err)
		fmt.Println(err)
		return fmt.Errorf("ERROR: %v", err)
	}

	return nil
}
