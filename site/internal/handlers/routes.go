package handlers

import (
	"net/http"
	"rapidart/internal/handlers/api"
	"rapidart/internal/handlers/middleware"
	"rapidart/internal/handlers/web"
)

const RES_ROUTE = "/res/"

type Middleware func(http.Handler) http.Handler
type Handler func(http.ResponseWriter, *http.Request)

type route struct {
	middlewares []Middleware
	handler     Handler
}

// UPDATE THIS WHEN NEW ADDING NEW ROUTE
var routes = map[string]route{
	/// WEB ROUTES
	"/": {
		[]Middleware{middleware.Auth},
		web.Index,
	},
	"/login/": {
		[]Middleware{},
		web.Login,
	},
	"/profile/": {
		[]Middleware{middleware.Auth},
		web.Profile,
	},
	"/drawing/": {
		[]Middleware{middleware.Auth},
		web.Drawing,
	},
	"/post/": {
		[]Middleware{middleware.Auth},
		web.Post,
	},
	"/search/": {
		[]Middleware{middleware.Auth},
		web.Search,
	},

	/// API ROUTES
	"/api/img/basiscanvas/": {
		[]Middleware{},
		api.BasisCanvas,
	},
	"/api/img/post/": {
		[]Middleware{middleware.Auth},
		api.GetPost,
	},
	"/api/save_post": {
		[]Middleware{middleware.Auth},
		api.SavePost,
	},
}
