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
	"GET /{$}": {
		[]Middleware{middleware.RequireAuth},
		web.Index,
	},
	"/login/": {
		[]Middleware{middleware.RequireNoAuth},
		web.Login,
	},
	"/register/": {
		[]Middleware{middleware.RequireNoAuth},
		web.Register,
	},
	"GET /profile/{username...}": {
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
	"/toplist/": {
		[]Middleware{middleware.RequireAuth},
		web.Toplist,
	},

	/// API ROUTES
	"/api/auth/login/": {
		[]Middleware{middleware.RequireNoAuth},
		api.AuthLogin,
	},
	"/api/auth/logout/": {
		[]Middleware{middleware.RequireAuth},
		api.AuthLogout,
	},
	"POST /api/user/register/": {
		[]Middleware{middleware.RequireNoAuth},
		api.UserRegister,
	},
	"/api/img/basiscanvas/": {
		[]Middleware{middleware.RequireAuth},
		api.ImgBasisCanvas,
	},
	"/api/img/post/": {
		[]Middleware{middleware.RequireAuth},
		api.GetPost,
	},
	"/api/save-post": {
		[]Middleware{middleware.RequireAuth},
		api.SavePost,
	},
	"/api/img/user/profile-pic/": {
		[]Middleware{middleware.RequireAuth},
		api.ImgUserProfilePic,
	},
}
