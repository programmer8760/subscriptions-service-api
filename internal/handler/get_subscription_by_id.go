package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/programmer8760/subscriptions-service-api/internal/domain"
)

// @Summary Get subscription by ID
// @Description Fetches a subscription with specified ID
// @Tags subscriptions
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} domain.Subscription
// @Failure 400 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /subscriptions/{id} [get]
func (h *Handler) GetSubscriptionByID(w http.ResponseWriter, r *http.Request) {
	id64, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		WriteErrorJSON(w, domain.ErrInvalidID, 400)
		return
	}
	ctx := r.Context()
	log := h.logger.With("request_id", ctx.Value(domain.RequestIDKey))

	resp, err := h.subscriptions.GetSubscriptionByID(ctx, uint(id64))
	if err != nil {
		if errors.Is(err, domain.ErrSubscriptionNotFound) {
			WriteErrorJSON(w, err, 404)
			return
		}
		if errors.Is(err, domain.ErrInvalidID) {
			WriteErrorJSON(w, err, 400)
			return
		}
		log.Error("failed to get subscription by id", "err", err)
		WriteErrorJSON(w, errors.New("internal server error"), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
