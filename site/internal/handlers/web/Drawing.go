package web

import (
	"log"
	"net/http"
	"rapidart/internal/database"
	"rapidart/internal/models"
	"rapidart/internal/util"
	"strconv"
)

type DrawingPageData struct {
	Title       string
	BasisCanvas models.BasisCanvas // Add BasisCanvas for the drawing page
}

func Drawing(w http.ResponseWriter, r *http.Request) {
	// Get the basis id parameter
	basisIDStr := r.URL.Query().Get("line")
	basisID, err := strconv.Atoi(basisIDStr) // Convert to int
	if err != nil || basisID <= 0 {
		http.Error(w, "Invalid basis canvas ID", http.StatusBadRequest)
		return
	}

	// Get the specific basiscanvas with API
	canvas, err := database.GetBasisCanvasById(basisID)
	if err != nil {
		log.Println("Error fetching basis canvas:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Prepare data to send to the templates
	pageData := DrawingPageData{
		Title:       "Drawing",
		BasisCanvas: canvas,
	}

	// Renders drawing.tmpl with template for basis canvases
	err = util.HttpServeTemplate("draw.tmpl", false, pageData, w)
	if err != nil {
		log.Println("Error serving template:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
