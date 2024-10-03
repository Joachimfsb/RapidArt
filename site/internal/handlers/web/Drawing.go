package web

import (
	"log"
	"net/http"
	"rapidart/internal/util"
)

func Drawing(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		drawingGetRequest(w, r)
	default:
		http.Error(w, "this method is not supported", http.StatusNotImplemented)
	}
}

func drawingGetRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello drawingGetRequest")

	var headerTitle = Title{
		Title: "Drawing",
	}

	err := util.HttpServeTemplate("draw.tmpl", headerTitle, w)
	if err != nil {
		log.Println(err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
