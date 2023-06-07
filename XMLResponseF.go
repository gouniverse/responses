package responses

import "net/http"

// XMLResponseF - responds with the result from a string returning function
func XMLResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	XMLResponse(w, r, s)
}
