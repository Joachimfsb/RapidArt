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
		[]Middleware{middleware.RequireAuth},
		web.Index,
	},
	"/login/": {
		[]Middleware{middleware.RequireNoAuth},
		web.Login,
	},
	"/profile/": {
		[]Middleware{middleware.RequireAuth},
		web.Profile,
	},
	"/drawing/": {
		[]Middleware{middleware.RequireAuth},
		web.Drawing,
	},
	"/post/": {
		[]Middleware{middleware.RequireAuth},
		web.Post,
	},
	"/search/": {
		[]Middleware{middleware.RequireAuth},
		web.Search,
	},

	/// API ROUTES
	"/api/auth/login/": {
		[]Middleware{middleware.RequireNoAuth},
		api.Login,
	},
	"/api/auth/logout/": {
		[]Middleware{middleware.RequireAuth},
		api.Logout,
	},
	"/api/img/basiscanvas/": {
		[]Middleware{middleware.RequireAuth},
		api.BasisCanvas,
	},
	"/api/img/post/": {
		[]Middleware{middleware.RequireAuth},
		api.GetPost,
	},
	"/api/save_post": {
		[]Middleware{middleware.RequireAuth},
		api.SavePost,
	},
}
