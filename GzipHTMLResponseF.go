package responses

import "net/http"

// GzipHTMLResponseF - responds with the result from a string returning function
func GzipHTMLResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	GzipHTMLResponse(w, r, s)
}
