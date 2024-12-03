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
func HttpServeTemplate(tmpl string, partial bool, model any, w http.ResponseWriter) error {
	// Are accessible to the templates (if many functions are added here, this map should be initialized once elsewhere)
	funcs := template.FuncMap{
		"add": func(i int, j int) int { return i + j },
	}

	// Determine path to partials file
	dir := ""
	if !partial {
		dir = glob.HTML_DIR
	} else {
		dir = glob.HTML_PARTIALS_DIR
	}
	tmplFiles := []string{
		filepath.Join(dir, tmpl),
		filepath.Join(glob.HTML_PARTIALS_DIR, "header.tmpl"), // Header is always available to templates
	}
	t, err := template.New(tmpl).Funcs(funcs).ParseFiles(tmplFiles...)
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

// Get session token
//
// Returns: Session token or empty string
func GetSessionTokenFromCookie(r *http.Request) string {

	//// Get currently logged in user ////
	// Get session cookie
	cookie, err := r.Cookie("session-token")
	if err != nil {
		return ""
	}

	return cookie.Value
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
