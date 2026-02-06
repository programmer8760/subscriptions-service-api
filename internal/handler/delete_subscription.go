package handler

import (
	"net/http"
	"strconv"

	"github.com/prajkin/em-test-task/internal/domain"
)

// @Summary Delete subscription
// @Description Deletes a subscription
// @Tags subscriptions
// @Param id path string true "Subscription ID"
// @Success 204 "No Content"
// @Failure 404 {string} string "subscription not found"
// @Failure 500 {string} string "internal server error"
// @Router /subscriptions/{id} [delete]
func (h *Handler) DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	id64, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	ctx := r.Context()

	err = h.subscriptions.DeleteSubscription(ctx, uint(id64))
	if err == domain.ErrSubscriptionNotFound {
		http.Error(w, err.Error(), 404)
		return
	} else if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
