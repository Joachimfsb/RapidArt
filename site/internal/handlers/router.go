package handlers

import (
	"log"
	"net/http"
	"rapidart/internal/consts"

	config "rapidart/internal/config"
)

var router *http.ServeMux // router

// Start the router (server)
func StartRouter() error {

	// Go middlewares: https://medium.com/geekculture/learn-go-middlewares-by-examples-da5dc4a3b9aa
	router = http.NewServeMux()

	// Set up routing
	serveStaticContent()
	bindRoutes() // Bind all routes

	// Start the server
	log.Println("Service is listening om port: " + config.Config.Server.Port)
	log.Fatal(http.ListenAndServe(config.Config.Server.Host+":"+config.Config.Server.Port, router))

	return nil
}

// Serve static content
func serveStaticContent() {
	// Resources
	router.Handle(RES_ROUTE, http.StripPrefix(RES_ROUTE, http.FileServer(http.Dir(consts.RES_DIR))))
}

// Bind the routes specified in routes.go
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
			router.Handle(url, current) // Handle
		} else {
			// No middlewares
			router.HandleFunc(url, route.handler) // Handle
		}
	}
}
