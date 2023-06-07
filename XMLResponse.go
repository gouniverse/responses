package responses

import "net/http"

// XMLResponse - responds with the result from a string returning function
func XMLResponse(w http.ResponseWriter, r *http.Request, body string) {
	w.Header().Set("Content-Type", "text/xml")
	w.Write([]byte(body))
}
