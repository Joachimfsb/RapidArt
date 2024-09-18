package util

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"rapidart/internal/glob"
	"strconv"
)

func HttpReturnError(status int, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Error "+strconv.Itoa(status))
}

// Parses and serves a template and a model to the http writer.
//
// ARG1: tmpl is the file path below globs.HTML_DIR. Example "index.tmpl"
func HttpServeTemplate(tmpl string, model any, w http.ResponseWriter) error {
	filep := filepath.Join(glob.HTML_DIR, tmpl)
	t, err := template.ParseFiles(filep)
	if err == nil {
		var buffer bytes.Buffer
		err = t.Execute(&buffer, model)
		if err == nil {
			buffer.WriteTo(w)
			return nil
		}
	}

	return errors.New("Something went wrong during parsing of template file " + tmpl + ". Got error [" + err.Error() + "]")
}

// Serves a single file to the writer.
//
// ARG1: file is the file path below globs.HTML_DIR. Example "index.html"
func HttpServeStatic(file string, w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(glob.HTML_DIR, file))
}
