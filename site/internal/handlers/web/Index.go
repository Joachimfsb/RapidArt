package web

import (
	"log"
	"net/http"
	"rapidart/internal/database"
	"rapidart/internal/models"
	"rapidart/internal/util"
	"time"
)

type IndexPageData struct {
	Title         string
	BasisCanvases []models.BasisCanvas
}

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		indexGetRequest(w, r)
	default:
		http.Error(w, "this method is not supported", http.StatusNotImplemented)
	}
}

func indexGetRequest(w http.ResponseWriter, r *http.Request) {

	// Get current time
	currentTime := time.Now()

	// Gets list of basis canvases based on current time
	canvases, err := database.GetBasisCanvasesByDateTime(currentTime)
	if err != nil {
		log.Println("Error fetching basis canvases:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Prepare the data to send to the template
	pageData := IndexPageData{
		Title:         "Index",
		BasisCanvases: canvases,
	}

	// Renders index.tmpl with template for basis canvases
	err = util.HttpServeTemplate("index.tmpl", pageData, w)
	if err != nil {
		log.Println("Error serving template:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
