package basismanager

import (
	"log"
	"rapidart/internal/database"
	"rapidart/internal/models"
	"time"
)

// Fetches basis canvases by the current time
func GetBasisCanvasesByDateTime(currentTime time.Time) ([]models.BasisCanvas, error) {
	canvases, err := database.GetBasisCanvasesByDateTime(currentTime)
	if err != nil {
		log.Println("Error fetching basis canvases:", err)
		return nil, err
	}
	return canvases, nil
}

func GetBasisCanvasById(basisCanvasId int) (models.BasisCanvas, error) {
	return database.GetBasisCanvasById(basisCanvasId)
}
