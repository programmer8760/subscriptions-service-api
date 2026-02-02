package handler

import (
	"net/http"
)

type Handler struct {
	routes *http.ServeMux
}

func NewHandler() *Handler {
	h := &Handler{
		routes: http.NewServeMux(),
	}

	// h.routes.HandleFunc("GET /subscriptions", h.GetSubscriptions)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.routes.ServeHTTP(w, r)
}
