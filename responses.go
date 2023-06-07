package responses

import (
	"compress/gzip"
	"net/http"
)

// GzipCSSResponse - responds with the CSS body
func GzipCSSResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "text/css")
	w.Header().Add("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	gz.Write([]byte(body))
	defer gz.Close()
}

// GzipCSSResponse - responds with the result from a string returning function
func GzipCSSResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	GzipCSSResponse(w, r, s)
}

// GzipJSResponse - responds with the JS body
func GzipJSResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "application/javascript")
	w.Header().Add("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	gz.Write([]byte(body))
	defer gz.Close()
}

// GzipJSResponse - responds with the result from a string returning function
func GzipJSResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	GzipJSResponse(w, r, s)
}

// GzipJSONResponse - responds with the JSON body
func GzipJSONResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	gz.Write([]byte(body))
	defer gz.Close()
}

// GzipJSONResponse - responds with the result from a string returning function
func GzipJSONResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	GzipJSONResponse(w, r, s)
}

// GzipXMLResponse - responds with the XML body
func GzipXMLResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "text/xml")
	w.Header().Add("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	gz.Write([]byte(body))
	defer gz.Close()
}

// GzipXMLResponse - responds with the result from a string returning function
func GzipXMLResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	GzipXMLResponse(w, r, s)
}
