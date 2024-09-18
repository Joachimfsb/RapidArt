package handlers

import (
	"net/http"
	"rapidart/internal/handlers/web"
)

const (
	RES_ROUTE = "/res/"
	RES_DIR   = "web/res/"
	HTML_DIR  = "web/html/"
)

var routes = map[string]func(http.ResponseWriter, *http.Request){
	/// WEB ROUTES
	"/": web.Index,

	/// API ROUTES
	// ...
}
