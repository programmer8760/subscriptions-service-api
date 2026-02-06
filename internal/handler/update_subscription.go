package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/prajkin/em-test-task/internal/domain"
	"github.com/prajkin/em-test-task/internal/dto"
)

type UpdateSubscriptionRequest struct {
	Name      *string      `json:"name,omitempty"`
	Price     *int         `json:"price,omitempty"`
	UserID    *uuid.UUID   `json:"user_id,omitempty"`
	StartDate *domain.Date `json:"start_date,omitempty"`
	EndDate   *domain.Date `json:"end_date,omitempty"`
}

// @Summary Update subscription
// @Description Updates an existing subscription
// @Tags subscriptions
// @Accept json
// @Param id path string true "Subscription ID"
// @Param request body UpdateSubscriptionRequest true "Subscription payload"
// @Success 204 "No Content"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /subscriptions [put]
func (h *Handler) UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	var req UpdateSubscriptionRequest
	id64, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	ctx := r.Context()

	err = h.subscriptions.UpdateSubscription(ctx, dto.UpdateSubscriptionDTO{
		ID:        uint(id64),
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

	w.WriteHeader(http.StatusNoContent)
}
