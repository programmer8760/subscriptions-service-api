package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/prajkin/em-test-task/internal/domain"
)

type CreateSubscriptionRequest struct {
	Name      string       `json:"name"`
	Price     int          `json:"price"`
	UserID    uuid.UUID    `json:"user_id"`
	StartDate domain.Date  `json:"start_date" example:"01-2006" swaggertype:"string"`
	EndDate   *domain.Date `json:"end_date,omitempty" example:"01-2006" swaggertype:"string"`
}

// @Summary Create subscription
// @Description Creates a new subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param request body CreateSubscriptionRequest true "Subscription payload"
// @Success 201 {object} domain.Subscription
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /subscriptions [post]
func (h *Handler) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var req CreateSubscriptionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteErrorJSON(w, err, 400)
		return
	}

	ctx := r.Context()
	h.logger.Info("request body", "request_id", ctx.Value(domain.RequestIDKey), "input", req)

	resp, err := h.subscriptions.CreateSubscription(ctx, domain.Subscription{
		Name:      req.Name,
		Price:     req.Price,
		UserID:    req.UserID,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	})
	if err != nil {
		var br domain.BadRequest
		if errors.As(err, &br) {
			WriteErrorJSON(w, err, 400)
			return
		}
		h.logger.Error("failed to create subscription", "request_id", ctx.Value(domain.RequestIDKey), "err", err)
		WriteErrorJSON(w, errors.New("internal server error"), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
