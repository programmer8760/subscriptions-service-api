package handler

import (
	"encoding/json"
	"net/http"
)

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
