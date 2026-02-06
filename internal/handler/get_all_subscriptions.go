package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/prajkin/em-test-task/internal/domain"
)

// @Summary Get all subscriptions
// @Description Returns a list of all subscriptions
// @Tags subscriptions
// @Accept json
// @Produce json
// @Success 200 {array} domain.Subscription
// @Failure 500 {object} domain.ErrorResponse
// @Router /subscriptions [get]
func (h *Handler) GetAllSubscriptions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp, err := h.subscriptions.GetAllSubscriptions(ctx)
	if err != nil {
		h.logger.Error("failed to get all subscriptions", "request_id", ctx.Value(domain.RequestIDKey), "err", err)
		WriteErrorJSON(w, errors.New("internal server error"), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
