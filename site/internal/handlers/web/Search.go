package web

import (
	"log"
	"net/http"
	"rapidart/internal/util"
)

func Search(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		searchGetRequest(w, r)
	default:
		http.Error(w, "This method is not supported", http.StatusNotImplemented)
	}
}

type Title struct {
	Title string
}

// searchGetRequest serves the search page template.
func searchGetRequest(w http.ResponseWriter, r *http.Request) {
	var headerTitle = Title{
		Title: "Search",
	}

	err := util.HttpServeTemplate("search.tmpl", headerTitle, w)
	if err != nil {
		log.Println("Error serving search template:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
