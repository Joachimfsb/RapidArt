package web

import (
	"log"
	"net/http"
	"rapidart/internal/util"
)

type Title struct {
	Title string
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		loginGetRequest(w, r)
	default:
		http.Error(w, "this method is not supported", http.StatusNotImplemented)
	}
}

func loginGetRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello LoginGetRequest")

	var headerTitle = Title{
		Title: "Log in",
	}

	err := util.HttpServeTemplate("login.tmpl", headerTitle, w)
	if err != nil {
		log.Println(err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
