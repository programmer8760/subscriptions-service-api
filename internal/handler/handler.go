package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/prajkin/em-test-task/internal/domain"
	"github.com/prajkin/em-test-task/internal/service"
	swagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	routes        *http.ServeMux
	subscriptions *service.SubscriptionsService
	logger        *slog.Logger
}

func NewHandler(subs *service.SubscriptionsService, log *slog.Logger) *Handler {
	h := &Handler{
		routes:        http.NewServeMux(),
		subscriptions: subs,
		logger:        log,
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
	start := time.Now()

	reqID := r.Header.Get("X-Request-ID")
	if reqID == "" {
		reqID = uuid.NewString()
	}

	ctx := context.WithValue(r.Context(), domain.RequestIDKey, reqID)
	r = r.WithContext(ctx)
	ww := NewWrapWriter(w)

	h.routes.ServeHTTP(ww, r)

	h.logger.Info("request",
		"request_id", reqID,
		"method", r.Method,
		"path", r.URL.Path,
		"query", r.URL.RawQuery,
		"status", ww.Status(),
		"duration_ms", time.Since(start).Milliseconds(),
	)
}

func WriteErrorJSON(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	resp := domain.ErrorResponse{Error: err.Error()}
	json.NewEncoder(w).Encode(resp)
}
