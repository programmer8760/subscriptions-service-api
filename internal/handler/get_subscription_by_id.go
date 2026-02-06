package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/prajkin/em-test-task/internal/domain"
)

// @Summary Get subscription by ID
// @Description Fetches a subscription with specified ID
// @Tags subscriptions
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} domain.Subscription
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

	resp, err := h.subscriptions.GetSubscriptionByID(ctx, uint(id64))
	if err == domain.ErrSubscriptionNotFound {
		WriteErrorJSON(w, err, 404)
		return
	} else if err != nil {
		WriteErrorJSON(w, err, 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
