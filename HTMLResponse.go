package responses

import "net/http"

// HTMLResponse - responds with the HTML body
func HTMLResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(body))
}
