package mhttp

import (
	"mime"
	"net/http"
)

func SetHTMLContentType(w http.ResponseWriter) {
	setContentType(w, "text/html")
}

func SetJSONContentType(w http.ResponseWriter) {
	setContentType(w, "application/json")
}

func SetFormContentType(w http.ResponseWriter) {
	setContentType(w, "x-www-form-urlencoded")
}

func SetPNGContentType(w http.ResponseWriter) {
	setContentType(w, "image/png")
}

func SetRSSContentType(w http.ResponseWriter) {
	setContentType(w, "application/rss+xml")
}

func SetAtomContentType(w http.ResponseWriter) {
	setContentType(w, "application/atom+xml")
}

func SetContentTypeFromExtension(w http.ResponseWriter, extension string) {
	mime := mime.TypeByExtension(extension)
	if mime != "" {
		setContentType(w, mime)
	}
}

func setContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
}
