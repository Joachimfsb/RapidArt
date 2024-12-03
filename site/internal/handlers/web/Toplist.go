package web

import (
	"log"
	"net/http"
	"rapidart/internal/basismanager"
	"rapidart/internal/models"
	"rapidart/internal/util"
	"time"
)

type ToplistPageData struct {
	BasisCanvases []models.BasisCanvas
}

func Toplist(w http.ResponseWriter, r *http.Request) {
	// Get current time
	currentTime := time.Now()

	// Fetch basis canvases
	canvases, err := basismanager.GetBasisCanvasesByDateTime(currentTime)
	if err != nil {
		log.Println("Error fetching basis canvases:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Prepare the data to send to the template
	pageData := ToplistPageData{
		BasisCanvases: canvases,
	}

	// Render toplist.tmpl with template for basis canvases, top posts, and top users
	err = util.HttpServeTemplate("toplist.tmpl", false, pageData, w)
	if err != nil {
		log.Println("Error serving template:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
