package http

import "net/http"

type HttpHandler struct {
}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello"))
}

func HandleCache(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte("Handle cache"))
}
