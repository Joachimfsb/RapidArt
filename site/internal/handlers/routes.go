package handlers

import (
	"net/http"
	"rapidart/internal/handlers/web"
)

const RES_ROUTE = "/res/"

var routes = map[string]func(http.ResponseWriter, *http.Request){
	/// WEB ROUTES
	"/": web.Index,

	/// API ROUTES
	// ...
}
