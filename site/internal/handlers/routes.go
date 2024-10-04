package handlers

import (
	"net/http"
	"rapidart/internal/handlers/api"
	"rapidart/internal/handlers/web"
)

const RES_ROUTE = "/res/"

var routes = map[string]func(http.ResponseWriter, *http.Request){
	/// WEB ROUTES
	"/":         web.Index,
	"/login/":   web.Login,
	"/profile/": web.Profile,
	"/drawing/": web.Drawing,
	"/post/":    web.Post,
	"/search/":  web.Search,

	/// API ROUTES
	"/api/img/basiscanvas/": api.BasisCanvas,
}
