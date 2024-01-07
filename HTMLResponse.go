package responses

import "net/http"

// HTMLResponse - responds with the HTML body
func HTMLResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(body))
}
