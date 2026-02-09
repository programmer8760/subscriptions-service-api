package service

import (
	"context"
	"strings"

	"github.com/programmer8760/subscriptions-service-api/internal/domain"
	"github.com/programmer8760/subscriptions-service-api/internal/dto"
)

func (s *subscriptionsService) UpdateSubscription(ctx context.Context, req dto.UpdateSubscriptionDTO) (domain.Subscription, error) {
	if req.ID == 0 {
		return domain.Subscription{}, domain.BadRequest{Err: domain.ErrInvalidID}
	}
	if req.Name == nil && req.Price == nil && req.UserID == nil && req.StartDate == nil && req.EndDate == nil {
		return domain.Subscription{}, domain.BadRequest{Err: domain.ErrNoChanges}
	}
	if req.Name != nil {
		if *req.Name = strings.TrimSpace(*req.Name); *req.Name == "" {
			return domain.Subscription{}, domain.BadRequest{Err: domain.ErrInvalidName}
		}
	}
	if req.Price != nil {
		if *req.Price <= 0 {
			return domain.Subscription{}, domain.BadRequest{Err: domain.ErrInvalidPrice}
		}
	}
	if req.StartDate != nil {
		if (*req.StartDate).Time.IsZero() {
			return domain.Subscription{}, domain.BadRequest{Err: domain.ErrInvalidStartDate}
		}
	}
	if req.EndDate != nil {
		if (*req.EndDate).Time.IsZero() {
			return domain.Subscription{}, domain.BadRequest{Err: domain.ErrInvalidEndDate}
		}
	}

	sub, err := s.repo.Update(ctx, req)
	if err != nil {
		return domain.Subscription{}, err
	}

	s.logger.Info("subscription updated", "request_id", ctx.Value(domain.RequestIDKey), "payload", req)
	return sub, nil
}
