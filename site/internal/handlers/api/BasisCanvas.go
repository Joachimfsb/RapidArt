package api

import (
	"net/http"
	"rapidart/internal/database"
	"rapidart/internal/util"
	"strconv"
)

// ////////////// HANDLER /////////////// //

// Img Basis canvas handler. This function routes the different REST methods to other handlers.
func ImgBasisCanvas(w http.ResponseWriter, r *http.Request) {

	// Check if id is specified
	if !r.URL.Query().Has("id") {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}
	// Try to convert to int
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		util.HttpReturnError(http.StatusBadRequest, w)
		return
	}

	// Fetch image by id from DB
	canvas, err := database.GetBasisCanvasById(id)
	if err != nil {
		util.HttpReturnError(http.StatusNotFound, w)
		return
	}

	// Return image
	w.Header().Set("Content-Type", "image/png")

	w.Write(canvas.Image)
}
