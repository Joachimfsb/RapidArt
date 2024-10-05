package database

import (
	"database/sql"
	"errors"
	"fmt"
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
			"WHERE ? BETWEEN BG.StartTimestamp AND BG.EndTimestamp",
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
