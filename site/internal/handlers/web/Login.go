package web

import (
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

	util.HttpServeStatic("login.html", w, r)
}
