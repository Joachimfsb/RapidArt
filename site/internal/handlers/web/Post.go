package web

import (
	"log"
	"net/http"
	"rapidart/internal/util"
)

type Title struct {
	Title string
}

func Post(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		postGetRequest(w, r)
	default:
		http.Error(w, "this method is not supported", http.StatusNotImplemented)
	}
}

func postGetRequest(w http.ResponseWriter, r *http.Request) {

	var headerTitle = Title{
		Title: "Post",
	}

	err := util.HttpServeTemplate("post.tmpl", headerTitle, w)
	if err != nil {
		log.Println(err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
