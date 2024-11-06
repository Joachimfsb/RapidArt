package util

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"rapidart/internal/glob"
	"strconv"
	"strings"
)

func HttpReturnError(status int, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Error "+strconv.Itoa(status))
}

// Parses and serves a template (with additionals (header)) and a model to the http writer.
//
// ARG1: tmpl is the file path below globs.HTML_DIR. Example "index.tmpl"
func HttpServeTemplate(tmpl string, model any, w http.ResponseWriter) error {
	tmplFiles := []string{
		filepath.Join(glob.HTML_DIR, tmpl),
		filepath.Join(glob.HTML_DIR, "header.tmpl"),
	}
	t, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		return fmt.Errorf("error parsing template files %v: %w", tmplFiles, err)
	}

	var buffer bytes.Buffer
	err = t.Execute(&buffer, model)
	if err != nil {
		return fmt.Errorf("error executing template %s: %w", tmpl, err)
	}

	buffer.WriteTo(w)
	return nil
}

// Serves a single file to the writer.
//
// ARG1: file is the file path below globs.HTML_DIR. Example "index.html"
func HttpServeStatic(file string, w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(glob.HTML_DIR, file))
}

// Simplifies user agent field
func UserAgentToBrowser(ua string) string {

	identifiers := map[string]string{
		"Chrome":  "Chrome",
		"Firefox": "Firefox",
		"Safari":  "Safari",
		"MSIE":    "Internet explorer",
		"Trident": "Internet explorer 11",
		"Edge":    "Edge",
		"Opera":   "Opera",
		"OPR":     "Opera",
	}

	for _, id := range identifiers {
		if strings.Contains(ua, id) {
			return id
		}
	}
	return "Unknown"
}
