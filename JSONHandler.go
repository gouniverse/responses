package responses

import (
	"net/http"
)

type JSONHandler func(w http.ResponseWriter, r *http.Request) string

func (h JSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	str := h(w, r)
	JSONResponse(w, r, str)
}
