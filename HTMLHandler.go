package responses

import (
	"net/http"
)

type HTMLHandler func(w http.ResponseWriter, r *http.Request) string

func (h HTMLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	str := h(w, r)
	HTMLResponse(w, r, str)
}
