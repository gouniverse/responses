package responses

import "net/http"

// CSSResponseF - responds with the result from a string returning function
func CSSResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	w.Header().Set("Content-Type", "text/css")
	w.Write([]byte(s))
}
