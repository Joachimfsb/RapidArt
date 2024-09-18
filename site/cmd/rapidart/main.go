package main

import (
	"log"
	"net/http"
	"os"
	"rapidart/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Setting default port 8080")
		port = "8080"
	}

	handlers.ServeStaticContent()
	handlers.BindRoutes() // Bind all routes

	log.Println("Service is listening om port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
