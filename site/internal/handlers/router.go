package handlers

import (
	"log"
	"net/http"
	"rapidart/internal/glob"
	"rapidart/internal/util"
)

var mux *http.ServeMux // router

func StartRouter() error {

	// Go middlewares: https://medium.com/geekculture/learn-go-middlewares-by-examples-da5dc4a3b9aa
	mux = http.NewServeMux()

	// Set up routing
	serveStaticContent()
	bindRoutes() // Bind all routes

	// Start the server
	log.Println("Service is listening om port: " + util.Config.Server.Port)
	log.Fatal(http.ListenAndServe(util.Config.Server.Host+":"+util.Config.Server.Port, mux))

	return nil
}

// TODO: Skru av directory listing:
// https://stackoverflow.com/questions/49589685/good-way-to-disable-directory-listing-with-http-fileserver-in-go
func serveStaticContent() {
	mux.Handle(RES_ROUTE, http.StripPrefix(RES_ROUTE, http.FileServer(http.Dir(glob.RES_DIR))))
}

func bindRoutes() {
	// Bind all the routes
	for url, route := range routes {
		// Chain middlewares in middleware list
		if len(route.middlewares) > 0 {
			current := route.middlewares[len(route.middlewares)-1](http.HandlerFunc(route.handler))
			// Reversed for loop
			for i := len(route.middlewares) - 2; i >= 0; i-- {
				current = route.middlewares[i](current)
			}
			mux.Handle(url, current) // Handle
		} else {
			// No middlewares
			mux.HandleFunc(url, route.handler) // Handle
		}
	}
}
