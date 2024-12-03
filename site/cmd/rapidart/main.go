package main

import (
	"log"
	"path/filepath"
	config "rapidart/internal/config"
	"rapidart/internal/consts"
	"rapidart/internal/database"
	"rapidart/internal/handlers"
)

func main() {
	// Initialize config
	err := config.InitializeConfig()
	if err != nil {
		log.Fatal("FATAL: Could not load config.json. Please check that the file '" + filepath.Join(consts.CONFIG_DIR, "config.json") + "' exists and is filled out properly. Error recieved: [" + err.Error() + "]")
	}
	log.Println("Config initialized")

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

	// Start server
	handlers.StartRouter()

}
