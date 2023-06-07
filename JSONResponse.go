package responses

import "net/http"

// JSONResponseF - responds with the JSON body
func JSONResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(body))
}
