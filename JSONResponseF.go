package responses

import "net/http"

// JSONResponseF - responds with the result from a string returning function
func JSONResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string) {
	s := f(w, r)
	JSONResponse(w, r, s)
}
