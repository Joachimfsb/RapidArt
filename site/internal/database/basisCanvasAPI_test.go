package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"rapidart/internal/models"
	"rapidart/test"
	"testing"
	"time"
)

func TestGetBasisCanvasByIdById(t *testing.T) {
	// Test data
	gallery := test.GenBasisGallery()
	canvas := test.GenBasisCanvas(gallery.BasisGalleryId)

	// Declare expectations
	mock.ExpectQuery("^SELECT").WithArgs(canvas.BasisCanvasId).WillReturnRows(GenRows([]models.BasisCanvas{canvas}))

	// Function call
	_, err := GetBasisCanvasById(canvas.BasisCanvasId)
	if err != nil {
		t.Fatal("Got error trying to get gallery by id: " + err.Error())
	}
	// Check if expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}

// Laget av KI/AI
// https://chatgpt.com/share/671fb9c4-80ec-800c-aa33-01d1926c539b
func TestGetBasisCanvasesByDateTime(t *testing.T) {
	// Test data
	gallery := test.GenBasisGallery()
	canvas1 := test.GenBasisCanvas(gallery.BasisGalleryId)
	canvas2 := test.GenBasisCanvas(gallery.BasisGalleryId)

	// Declare expectations with only one argument (timestamp)
	timestamp := time.Now().Local()
	mock.ExpectQuery("^SELECT").
		WithArgs(timestamp).
		WillReturnRows(GenRows([]models.BasisCanvas{canvas1, canvas2}))

	// Function call
	canvases, err := GetBasisCanvasesByDateTime(timestamp)
	if err != nil {
		t.Fatal("Got error trying to get canvases by time: " + err.Error())
	}

	if len(canvases) <= 0 {
		t.Fatal("Couldn't get any canvases: " + err.Error())
	}

	// Check if expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}

func TestAddNewCanvas(t *testing.T) {

	// Declare expectations
	//mock.ExpectCommit()
	gallery := test.GenBasisGallery()
	canvas := test.GenBasisCanvas(gallery.BasisGalleryId)

	//mock.ExpectBegin()
	mock.ExpectExec(`^INSERT (.+)`).WithArgs(canvas.BasisGalleryId, canvas.Type, canvas.Image).WillReturnResult(sqlmock.NewResult(1, 1))

	// Function call
	if err := AddNewCanvas(canvas); err != nil {
		t.Fatal("Got error trying to add report: " + err.Error())
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}
