package mhttp

import (
	"mime"
	"net/http"
)

const (
	MIMETypeText = "text/plain"
	MIMETypeHTML = "text/html"
	MIMETypePNG  = "image/png"
	MIMETypeForm = "x-www-form-urlencoded"
	MIMETypeJSON = "application/json"
	MIMETypeRSS  = "application/rss+xml"
	MIMETypeAtom = "application/atom+xml"
)

// SetResponseContentType sets response Content-Type header to a specified value.
func SetResponseContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
}

// SetResponseContentTypeFromExtension sets response Content-Type header from a file extension.
func SetResponseContentTypeFromExtension(w http.ResponseWriter, extension string) {
	mime := mime.TypeByExtension(extension)
	if mime != "" {
		setContentType(w, mime)
	}
}
