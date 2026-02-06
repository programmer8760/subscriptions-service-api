package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/prajkin/em-test-task/internal/domain"
	"github.com/prajkin/em-test-task/internal/dto"
)

type GetTotalPriceResponse struct {
	Price int `json:"price"`
}

// @Summary Get total subscriptions price
// @Description Counts total price of subscriptions for specified period
// @Tags subscriptions
// @Produce json
// @Param from query string true "Start of period"
// @Param to query string true "End of period"
// @Param user_id query string false "User ID to filter subscriptions by"
// @Param name query string false "Name of subscription to filter by"
// @Success 200 {object} GetTotalPriceResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /subscriptions/price [get]
func (h *Handler) GetTotalPrice(w http.ResponseWriter, r *http.Request) {
	from, err := time.Parse("01-2006", r.URL.Query().Get("from"))
	if err != nil {
		WriteErrorJSON(w, domain.ErrInvalidFromDate, 400)
		return
	}
	to, err := time.Parse("01-2006", r.URL.Query().Get("to"))
	if err != nil {
		WriteErrorJSON(w, domain.ErrInvalidToDate, 400)
		return
	}
	var name *string
	if n := r.URL.Query().Get("name"); n != "" {
		name = &n
	}
	var userID *uuid.UUID
	if uidStr := r.URL.Query().Get("user_id"); uidStr != "" {
		if uid, err := uuid.Parse(uidStr); err != nil {
			WriteErrorJSON(w, err, 400)
			return
		} else {
			userID = &uid
		}
	}

	ctx := r.Context()
	h.logger.Info("request body", "request_id", ctx.Value(domain.RequestIDKey), "from", from, "to", to, "name", name, "user_id", userID)

	resp, err := h.subscriptions.GetTotalPrice(ctx, dto.GetTotalPriceDTO{
		From:   domain.NewDate(from),
		To:     domain.NewDate(to),
		Name:   name,
		UserID: userID,
	})
	if err != nil {
		var br domain.BadRequest
		if errors.As(err, &br) {
			WriteErrorJSON(w, err, 400)
			return
		}
		h.logger.Error("failed to get total subscriptions price", "request_id", ctx.Value(domain.RequestIDKey), "err", err)
		WriteErrorJSON(w, errors.New("internal server error"), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GetTotalPriceResponse{Price: resp})
}
