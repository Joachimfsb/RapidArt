package web

import (
	"net/http"
	"rapidart/internal/util"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache") // Tell browser not to cache
	util.HttpServeStatic("login.html", w, r)
}
