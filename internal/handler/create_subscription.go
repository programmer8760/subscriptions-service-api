package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/prajkin/em-test-task/internal/domain"
)

type CreateSubscriptionRequest struct {
	Name      string       `json:"name"`
	Price     int          `json:"price"`
	UserID    uuid.UUID    `json:"user_id"`
	StartDate domain.Date  `json:"start_date"`
	EndDate   *domain.Date `json:"end_date,omitempty"`
}

// @Summary Create subscription
// @Description Creates a new subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param request body CreateSubscriptionRequest true "Subscription payload"
// @Success 201 {object} domain.Subscription
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /subscriptions [post]
func (h *Handler) CreateSubscription(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
