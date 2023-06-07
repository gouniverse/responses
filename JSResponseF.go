package responses

import "net/http"

// JSResponseF - responds with the result from a string returning function
func JSResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	JSResponse(w, r, s)
}
