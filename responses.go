package responses

import (
	"compress/gzip"
	"net/http"
)

// CSSResponse - responds with the CSS body
func CSSResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "text/css")
	w.Write([]byte(body))
}

// CSSResponseF - responds with the result from a string returning function
func CSSResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	w.Header().Set("Content-Type", "text/css")
	w.Write([]byte(s))
}

// HTMLResponse - responds with the HTML body
func HTMLResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(body))
}

// HTMLResponseF - responds with the result from a string returning function
func HTMLResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	HTMLResponse(w, r, s)
}

// JSResponse - responds with the JS body
func JSResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "application/javascript")
	w.Write([]byte(body))
}

// JSResponseF - responds with the result from a string returning function
func JSResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	JSResponse(w, r, s)
}

// JSONResponseF - responds with the JSON body
func JSONResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(body))
}

// JSONResponseF - responds with the result from a string returning function
func JSONResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	JSONResponse(w, r, s)
}

// XMLResponse - responds with the result from a string returning function
func XMLResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "text/xml")
	w.Write([]byte(body))
}

// XMLResponseF - responds with the result from a string returning function
func XMLResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	XMLResponse(w, r, s)
}

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
	GzipJSResponse(w, r, s)
}

// GzipHTMLResponse - responds with the HTML body
func GzipHTMLResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Add("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	gz.Write([]byte(body))
	defer gz.Close()
}

// GzipHTMLResponse - responds with the result from a string returning function
func GzipHTMLResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	GzipHTMLResponse(w, r, s)
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
