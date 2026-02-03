package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetAllSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	resp, err := h.subscriptions.GetAllSubscriptions(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
