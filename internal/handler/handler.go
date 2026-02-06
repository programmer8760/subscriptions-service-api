package handler

import (
	"net/http"

	"github.com/prajkin/em-test-task/internal/service"
	swagger "github.com/swaggo/http-swagger"
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

	h.routes.HandleFunc("GET /subscriptions", h.GetAllSubscriptions)
	h.routes.HandleFunc("GET /subscriptions/{id}", h.GetSubscriptionByID)
	h.routes.HandleFunc("POST /subscriptions", h.CreateSubscription)
	h.routes.HandleFunc("PUT /subscriptions/{id}", h.UpdateSubscription)
	h.routes.HandleFunc("DELETE /subscriptions/{id}", h.DeleteSubscription)
	h.routes.HandleFunc("GET /subscriptions/price", h.GetTotalPrice)

	h.routes.Handle("GET /swagger/", swagger.WrapHandler)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.routes.ServeHTTP(w, r)
}
