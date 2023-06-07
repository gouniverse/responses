package responses

import (
	"compress/gzip"
	"net/http"
)

// GzipHTMLResponse - responds with the HTML body
func GzipHTMLResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Add("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	gz.Write([]byte(body))
	defer gz.Close()
}
