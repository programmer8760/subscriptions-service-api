package handler

import (
	"encoding/json"
	"net/http"
)

// @Summary Get all subscriptions
// @Description Returns a list of all subscriptions
// @Tags subscriptions
// @Accept json
// @Produce json
// @Success 200 {array} domain.Subscription
// @Failure 500 {string} string "internal server error"
// @Router /subscriptions [get]
func (h *Handler) GetAllSubscriptions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp, err := h.subscriptions.GetAllSubscriptions(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
