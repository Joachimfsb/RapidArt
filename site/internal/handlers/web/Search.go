package web

import (
	"log"
	"net/http"
	"rapidart/internal/util"
)

// Serves the search page template.
func Search(w http.ResponseWriter, _ *http.Request) {
	err := util.HttpServeTemplate("search.tmpl", nil, w)
	if err != nil {
		log.Println("Error serving search template:", err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
