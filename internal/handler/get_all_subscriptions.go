package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/programmer8760/subscriptions-service-api/internal/domain"
	"github.com/programmer8760/subscriptions-service-api/internal/dto"
)

// @Summary Get all subscriptions
// @Description Returns a list of all subscriptions
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param page query string false "Page number" default(1)
// @Param page_size query string false "Number of subscriptions per page" default(20)
// @Success 200 {array} domain.Subscription
// @Failure 500 {object} domain.ErrorResponse
// @Router /subscriptions [get]
func (h *Handler) GetAllSubscriptions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.GetAllSubscriptionsDTO
	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		page64, err := strconv.ParseUint(pageStr, 10, 64)
		if err != nil {
			log.Debug("error parsing page number", "err", err)
			WriteErrorJSON(w, domain.ErrInvalidPage, 400)
			return
		}
		page := uint(page64)
		req.Page = &page
	}
	if pageSizeStr := r.URL.Query().Get("page_size"); pageSizeStr != "" {
		pageSize64, err := strconv.ParseUint(pageSizeStr, 10, 64)
		if err != nil {
			log.Debug("error parsing page size", "err", err)
			WriteErrorJSON(w, domain.ErrInvalidPageSize, 400)
			return
		}
		pageSize := uint(pageSize64)
		req.PageSize = &pageSize
	}

	resp, err := h.subscriptions.GetAllSubscriptions(ctx, req)
	if err != nil {
		h.logger.Error("failed to get all subscriptions", "request_id", ctx.Value(domain.RequestIDKey), "err", err)
		WriteErrorJSON(w, errors.New("internal server error"), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
