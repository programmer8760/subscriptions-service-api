package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/prajkin/em-test-task/internal/domain"
)

type CreateSubscriptionRequest struct {
	Name      string     `json:"name"`
	Price     int        `json:"price"`
	UserID    uuid.UUID  `json:"user_id"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}

func (h *Handler) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req CreateSubscriptionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	ctx := r.Context()

	resp, err := h.subscriptions.CreateSubscription(ctx, domain.Subscription{
		Name:      req.Name,
		Price:     req.Price,
		UserID:    req.UserID,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
