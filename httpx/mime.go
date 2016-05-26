package httpx

import (
	"mime"
	"net/http"
)

// SetResponseContentType sets response Content-Type header to a specified value.
func SetResponseContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
}

// SetResponseContentTypeFromExtension sets response Content-Type header from a file extension.
func SetResponseContentTypeFromExtension(w http.ResponseWriter, extension string) {
	mime := mime.TypeByExtension(extension)
	if mime != "" {
		SetResponseContentType(w, mime)
	}
}
