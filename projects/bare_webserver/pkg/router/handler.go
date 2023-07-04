package router

import (
	"encoding/json"
	"net/http"
)

type Handler func(r *http.Request) (statusCode int, data map[string]interface{})

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	statusCode, data := h(r)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
