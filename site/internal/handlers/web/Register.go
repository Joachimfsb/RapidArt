package web

import (
	"net/http"
	"rapidart/internal/util"
)

func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		registerGetRequest(w, r)
	default:
		http.Error(w, "this method is not supported", http.StatusNotImplemented)
	}
}

func registerGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache") // Tell browser not to cache
	util.HttpServeStatic("register.html", w, r)
}
