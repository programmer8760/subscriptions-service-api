package service

import (
	"context"
	"strings"

	"github.com/prajkin/em-test-task/internal/domain"
)

func (s *SubscriptionsService) CreateSubscription(ctx context.Context, sub domain.Subscription) (domain.Subscription, error) {
	if sub.Name = strings.TrimSpace(sub.Name); sub.Name == "" {
		return domain.Subscription{}, domain.ErrInvalidName
	}
	if sub.Price <= 0 {
		return domain.Subscription{}, domain.ErrInvalidPrice
	}
	if sub.StartDate.Time.IsZero() {
		return domain.Subscription{}, domain.ErrInvalidStartDate
	}

	if err := s.repo.Create(ctx, &sub); err != nil {
		return domain.Subscription{}, err
	}

	return sub, nil
}
