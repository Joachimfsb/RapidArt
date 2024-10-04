package database

import (
	"database/sql"
	"errors"
	"fmt"
	"rapidart/internal/models"
)

// Fetches BasisCanvas data based on the given ID
func GetBasisCanvasById(id int) (models.BasisCanvas, error) {
	var canvas models.BasisCanvas

	// Query DB
	row := db.QueryRow("SELECT BasisCanvasId, BasisGalleryId, Type, Image FROM `BasisCanvas` WHERE BasisCanvasId = ?", id)

	err := row.Scan(&canvas.BasisCanvasId, &canvas.BasisGalleryId, &canvas.Type, &canvas.Image)
	// No rows returned
	if errors.Is(err, sql.ErrNoRows) {
		return models.BasisCanvas{}, fmt.Errorf("no basisgallery found by that id")
	}
	// General error
	if err != nil {
		return models.BasisCanvas{}, err
	}

	return canvas, nil
}
