package handlers

import (
	"net/http"
	"rapidart/internal/handlers/api"
	"rapidart/internal/handlers/middleware"
	"rapidart/internal/handlers/web"
	"rapidart/internal/handlers/web/components"
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
	"GET /login/{$}": {
		[]Middleware{middleware.RequireNoAuth},
		web.Login,
	},
	"GET /register/{$}": {
		[]Middleware{middleware.RequireNoAuth},
		web.Register,
	},
	"GET /profile/{username...}": {
		[]Middleware{middleware.RequireAuth},
		web.Profile,
	},
	"GET /drawing/{$}": {
		[]Middleware{middleware.RequireAuth},
		web.Drawing,
	},
	"GET /post/{post_id}": {
		[]Middleware{middleware.RequireAuth},
		web.Post,
	},
	"GET /search/{$}": {
		[]Middleware{middleware.RequireAuth},
		web.Search,
	},
	"GET /toplist/{$}": {
		[]Middleware{middleware.RequireAuth},
		web.Toplist,
	},

	/// Web components
	"GET /top/{type}": {
		[]Middleware{middleware.RequireAuth},
		components.Top,
	},
	"GET /post/comments/{post_id}": {
		[]Middleware{middleware.RequireAuth},
		components.Comments,
	},

	/// API ROUTES
	"POST /api/auth/login/{$}": {
		[]Middleware{middleware.RequireNoAuth},
		api.AuthLogin,
	},
	"POST /api/auth/logout/{$}": {
		[]Middleware{middleware.RequireAuth},
		api.AuthLogout,
	},
	"POST /api/user/register/{$}": {
		[]Middleware{middleware.RequireNoAuth},
		api.UserRegister,
	},
	"POST /api/user/follow/{UserId}/{Value}": {
		[]Middleware{middleware.RequireAuth},
		api.UserFollow,
	},
	"GET /api/img/basiscanvas/": {
		[]Middleware{middleware.RequireAuth},
		api.ImgBasisCanvas,
	},
	"GET /api/img/post/{$}": {
		[]Middleware{middleware.RequireAuth},
		api.GetPost,
	},
	"POST /api/post/comment/{id}": {
		[]Middleware{middleware.RequireAuth},
		api.PostComment,
	},
	"POST /api/post/like/{id}": {
		[]Middleware{middleware.RequireAuth},
		api.PostLike,
	},
	"POST /api/post/unlike/{id}": {
		[]Middleware{middleware.RequireAuth},
		api.PostUnlike,
	},
	"POST /api/save-post": {
		[]Middleware{middleware.RequireAuth},
		api.SavePost,
	},
	"GET /api/img/user/profile-pic/": {
		[]Middleware{middleware.RequireAuth},
		api.ImgUserProfilePic,
	},
	"GET /api/search/users/": {
		[]Middleware{middleware.RequireAuth},
		api.SearchUsers,
	},
	"POST /api/post/report/{id}": {
		[]Middleware{middleware.RequireAuth},
		api.PostReport,
	},
}
