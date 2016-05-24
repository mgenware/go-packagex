package mhttp

import (
	"mime"
	"net/http"
)

// SetHTMLContentType sets response Content-Type header to "text/html".
func SetHTMLContentType(w http.ResponseWriter) {
	setContentType(w, "text/html")
}

// SetJSONContentType sets response Content-Type header to "application/json".
func SetJSONContentType(w http.ResponseWriter) {
	setContentType(w, "application/json")
}

// SetFormContentType sets response Content-Type header to "x-www-form-urlencoded".
func SetFormContentType(w http.ResponseWriter) {
	setContentType(w, "x-www-form-urlencoded")
}

// SetPNGContentType sets response Content-Type header to "image/png".
func SetPNGContentType(w http.ResponseWriter) {
	setContentType(w, "image/png")
}

// SetRSSContentType sets response Content-Type header to "application/rss+xml".
func SetRSSContentType(w http.ResponseWriter) {
	setContentType(w, "application/rss+xml")
}

// SetAtomContentType sets response Content-Type header to "application/atom+xml".
func SetAtomContentType(w http.ResponseWriter) {
	setContentType(w, "application/atom+xml")
}

// SetContentTypeFromExtension sets response Content-Type header from a user specified file extension.
func SetContentTypeFromExtension(w http.ResponseWriter, extension string) {
	mime := mime.TypeByExtension(extension)
	if mime != "" {
		setContentType(w, mime)
	}
}

func setContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
}
