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

	/*defer func() {
		err := database.CloseDatabase()
		if err != nil {
			log.Fatal("unable to close the db connection")
		} else {
			log.Println("closed database connection")
		}
	}()

	err := database.InitializeDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize the database connection: %v", err)
	} else {
		log.Println("database successfully initialized")
	}*/
	http.HandleFunc("/test/", handlers.TestHandler)

	log.Println("Service is listening om port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
