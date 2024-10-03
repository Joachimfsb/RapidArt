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
		http.Error(w, "this method is not supported", http.StatusNotImplemented)
	}
}

func searchGetRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello searchGetRequest")

	var headerTitle = Title{
		Title: "Search",
	}

	err := util.HttpServeTemplate("search.tmpl", headerTitle, w)
	if err != nil {
		log.Println(err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
