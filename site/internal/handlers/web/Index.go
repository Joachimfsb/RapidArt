package web

import (
	"log"
	"net/http"
	"rapidart/internal/util"
)

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		indexGetRequest(w, r)
	default:
		http.Error(w, "this method is not supported", http.StatusNotImplemented)
	}
}

func indexGetRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello indexGetRequest")

	var headerTitle = Title{
		Title: "Index",
	}

	err := util.HttpServeTemplate("index.tmpl", headerTitle, w)
	if err != nil {
		log.Println(err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
