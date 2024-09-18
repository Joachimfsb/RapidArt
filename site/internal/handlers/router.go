package handlers

import (
	"net/http"
)

// TODO: Skru av directory listing:
// https://stackoverflow.com/questions/49589685/good-way-to-disable-directory-listing-with-http-fileserver-in-go
func ServeStaticContent() {
	http.Handle(RES_ROUTE, http.StripPrefix(RES_ROUTE, http.FileServer(http.Dir(RES_DIR))))
}

func BindRoutes() {
	// Bind all the routes
	for url, handler := range routes {
		http.HandleFunc(url, handler)
	}
}
