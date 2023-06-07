package responses

import "net/http"

// CSSResponse - responds with the CSS body
func CSSResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "text/css")
	w.Write([]byte(body))
}
