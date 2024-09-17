package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Setting default port 8080")
		port = "8080"
	}

	log.Println("Service is listening om port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
