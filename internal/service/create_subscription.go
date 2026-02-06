package service

import (
	"context"
	"strings"

	"github.com/prajkin/em-test-task/internal/domain"
)

func (s *SubscriptionsService) CreateSubscription(ctx context.Context, sub domain.Subscription) (domain.Subscription, error) {
	if sub.Name = strings.TrimSpace(sub.Name); sub.Name == "" {
		return domain.Subscription{}, domain.BadRequest{Err: domain.ErrInvalidName}
	}
	if sub.Price <= 0 {
		return domain.Subscription{}, domain.BadRequest{Err: domain.ErrInvalidPrice}
	}
	if sub.StartDate.Time.IsZero() {
		return domain.Subscription{}, domain.BadRequest{Err: domain.ErrInvalidStartDate}
	}
	if sub.EndDate != nil {
		if (*sub.EndDate).Time.IsZero() {
			return domain.Subscription{}, domain.BadRequest{Err: domain.ErrInvalidEndDate}
		}
		if (*sub.EndDate).Time.Before(sub.StartDate.Time) {
			return domain.Subscription{}, domain.BadRequest{Err: domain.ErrEndBeforeStart}
		}
	}

	if err := s.repo.Create(ctx, &sub); err != nil {
		return domain.Subscription{}, err
	}

	s.logger.Info("subscription created", "request_id", ctx.Value(domain.RequestIDKey), "subscription", sub)
	return sub, nil
}
