package handler

import (
	"encoding/json"
	"errors"
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
	StartDate *domain.Date `json:"start_date,omitempty" example:"01-2006" swaggertype:"string"`
	EndDate   *domain.Date `json:"end_date,omitempty" example:"01-2006" swaggertype:"string"`
}

// @Summary Update subscription
// @Description Updates an existing subscription
// @Tags subscriptions
// @Accept json
// @Param id path string true "Subscription ID"
// @Param request body UpdateSubscriptionRequest true "Subscription payload"
// @Success 204 "No Content"
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /subscriptions/{id} [put]
func (h *Handler) UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	var req UpdateSubscriptionRequest
	id64, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		WriteErrorJSON(w, err, 400)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteErrorJSON(w, err, 400)
		return
	}

	ctx := r.Context()
	h.logger.Info("request body", "request_id", ctx.Value(domain.RequestIDKey), "input", req)

	err = h.subscriptions.UpdateSubscription(ctx, dto.UpdateSubscriptionDTO{
		ID:        uint(id64),
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
		if err == domain.ErrSubscriptionNotFound {
			WriteErrorJSON(w, err, 404)
			return
		}
		h.logger.Error("failed to update subscription", "request_id", ctx.Value(domain.RequestIDKey), "err", err)
		WriteErrorJSON(w, errors.New("internal server error"), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
