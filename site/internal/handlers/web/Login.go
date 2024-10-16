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
	w.Header().Set("Cache-Control", "no-cache") // Tell browser not to cache
	util.HttpServeStatic("login.html", w, r)
}
