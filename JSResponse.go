package responses

import "net/http"

// JSResponse - responds with the JS body
func JSResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "application/javascript")
	w.Write([]byte(body))
}
