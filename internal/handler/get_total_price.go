package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/prajkin/em-test-task/internal/domain"
	"github.com/prajkin/em-test-task/internal/dto"
)

func (h *Handler) GetTotalPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	from, err := time.Parse("01-2006", r.URL.Query().Get("from"))
	if err != nil {
		http.Error(w, domain.ErrInvalidFromDate.Error(), 400)
		return
	}
	to, err := time.Parse("01-2006", r.URL.Query().Get("to"))
	if err != nil {
		http.Error(w, domain.ErrInvalidToDate.Error(), 400)
		return
	}
	var name *string
	if n := r.URL.Query().Get("name"); n != "" {
		name = &n
	}
	var userID *uuid.UUID
	if uidStr := r.URL.Query().Get("user_id"); uidStr != "" {
		if uid, err := uuid.Parse(uidStr); err != nil {
			http.Error(w, err.Error(), 400)
			return
		} else {
			userID = &uid
		}
	}

	ctx := r.Context()

	resp, err := h.subscriptions.GetTotalPrice(ctx, dto.GetTotalPriceDTO{
		From:   domain.NewDate(from),
		To:     domain.NewDate(to),
		Name:   name,
		UserID: userID,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
