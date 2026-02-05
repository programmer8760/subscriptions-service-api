package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/prajkin/em-test-task/internal/domain"
)

func (h *Handler) GetSubscriptionByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id64, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, domain.ErrInvalidID.Error(), 400)
		return
	}
	ctx := r.Context()

	resp, err := h.subscriptions.GetSubscriptionByID(ctx, uint(id64))
	if err == domain.ErrSubscriptionNotFound {
		http.Error(w, err.Error(), 404)
		return
	} else if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
