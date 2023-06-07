package responses

import "net/http"

// HTMLResponseF - responds with the result from a string returning function
func HTMLResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	HTMLResponse(w, r, s)
}
