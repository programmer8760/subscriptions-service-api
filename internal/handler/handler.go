package handler

import (
	"net/http"

	"github.com/prajkin/em-test-task/internal/service"
)

type Handler struct {
	routes        *http.ServeMux
	subscriptions *service.SubscriptionsService
}

func NewHandler(subs *service.SubscriptionsService) *Handler {
	h := &Handler{
		routes:        http.NewServeMux(),
		subscriptions: subs,
	}

	// h.routes.HandleFunc("GET /subscriptions", h.GetSubscriptions)
	h.routes.HandleFunc("GET /subscriptions/{id}", h.GetSubscriptionByID)
	h.routes.HandleFunc("POST /subscriptions", h.CreateSubscription)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.routes.ServeHTTP(w, r)
}
