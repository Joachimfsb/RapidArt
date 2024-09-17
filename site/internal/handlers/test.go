package handlers

import "net/http"

// Registration-HANDLER:
func TestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method { //a switch for the supported methods. Only GET method is supported
	case http.MethodGet:
		TestGetRequest(w, r)
	default: //Error message if GET method is not used
		http.Error(w, "This method is not supported! Only POST, GET, PUT and DELETE are supported", http.StatusNotImplemented)
	}
}

func TestGetRequest(w http.ResponseWriter, r *http.Request) {
	println("Hello Test Handler")
}
