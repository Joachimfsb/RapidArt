package web

import (
	"html/template"
	"net/http"
	"path/filepath"
	"rapidart/internal/glob"
)

// /////////// TEMPLATE MODEL ////////////// //
type User struct {
	Name string
	Age  int
}

// ////////////// HANDLER /////////////// //

// Web index handler. This function routes the different REST methods to other handlers.
func Index(w http.ResponseWriter, r *http.Request) {

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
	t, _ := template.ParseFiles(filepath.Join(glob.HTML_DIR, "index.tmpl"))
	t.Execute(w, model)
	// Server statically
	//http.ServeFile(w, r, filepath.Join("web/html", "index.tmpl"))

}
