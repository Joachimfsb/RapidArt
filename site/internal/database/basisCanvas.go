package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"rapidart/internal/glob"
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
	var count = 0
	rows, err := db.Query("SELECT BasisGalleryId FROM Basisgallery")
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf(glob.NoGallery)
	}
	defer rows.Close()

	for rows.Next() {
		var id models.BasisGallery

		err = rows.Scan(&id.BasisGalleryId)
		if err != nil {
			log.Println(glob.ScanFailed)
			return err
		}
		count = count + 1
	}
	newBasisCanvas.BasisGalleryId = count
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

/*
EXAMPLE
	// Specify the relative file name
	fileName := "tmp.png" // Adjust as necessary

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Construct the relative path
	tempPicPath := filepath.Join(cwd, "internal", "database", fileName) //path to temporary picture

	profilePic, err := ioutil.ReadFile(tempPicPath)
	if err != nil {
		log.Println(glob.PictureNotFound)
	}

	test := models.BasisCanvas{
		Type:  "Basis",
		Image: profilePic,
	}

	database.AddNewCanvas(test)
*/
