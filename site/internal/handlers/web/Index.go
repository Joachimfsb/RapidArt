package web

import (
	"html/template"
	"net/http"
	"path/filepath"
)

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
	// Serve using template
	t, _ := template.ParseFiles(filepath.Join("web/html", "index.tmpl"))
	t.Execute(w, nil)
	// Server statically
	//http.ServeFile(w, r, filepath.Join("web/html", "index.tmpl"))

}
