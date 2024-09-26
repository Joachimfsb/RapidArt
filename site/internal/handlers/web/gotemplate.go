package web

import (
	"log"
	"net/http"
	"rapidart/internal/util"
)

// /////////// TEMPLATE MODEL ////////////// //
type User struct {
	Name string
	Age  int
}

// ////////////// HANDLER /////////////// //

// Web index handler. This function routes the different REST methods to other handlers.
func Template(w http.ResponseWriter, r *http.Request) {

	switch r.Method { //a switch for the supported methods. Only GET method is supported
	case http.MethodGet:
		get(w, r)
	default: //Error message if GET method is not used
		http.Error(w, "This method is not supported.", http.StatusNotImplemented)
	}
}

// Internal get handler for this route
func get(w http.ResponseWriter, r *http.Request) {
	// Fetch data and put into model
	// ...

	var model = User{
		Name: "Bob",
		Age:  30,
	}

	// Serve using template
	err := util.HttpServeTemplate("gotemplate.tmpl", model, w)
	if err != nil {
		log.Println(err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}

	// Server static html file
	//util.HttpServeStatic("index.html", w, r)
}
