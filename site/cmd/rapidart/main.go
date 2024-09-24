package main

import (
	"log"
	"net/http"
	"rapidart/internal/handlers"
	"rapidart/internal/util"
)

func main() {
	// Initialize config
	err := util.InitializeConfig()
	if err != nil {
		log.Fatal("FAIL! Could not initialize config, got error [" + err.Error() + "]")
	}
	log.Println("Config initialized")

	BAD CODE

	/*
		// Initialize database
		defer func() {
			err := database.CloseDatabase()
			if err != nil {
				log.Fatal("Unable to close the database connection")
			} else {
				log.Println("Closed database connection")
			}
		}()

		dbError := database.InitializeDatabase()
		if dbError != nil {
			log.Fatalf("Failed to initialize the database connection: %v", dbError)
		} else {
			log.Println("Database initialized")
		}
	*/

	// Set up routing
	handlers.ServeStaticContent()
	handlers.BindRoutes() // Bind all routes

	// Start the server
	log.Println("Service is listening om port: " + util.Config.Server.Port)
	log.Fatal(http.ListenAndServe(util.Config.Server.Host+":"+util.Config.Server.Port, nil))

}
