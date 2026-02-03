package service

import (
	"context"

	"github.com/prajkin/em-test-task/internal/domain"
)

func (s *SubscriptionsService) GetSubscriptionByID(ctx context.Context, id uint) (domain.Subscription, error) {
	if id == 0 {
		return domain.Subscription{}, domain.ErrInvalidID
	}

	sub, err := s.repo.GetByID(ctx, id)

	return sub, err
}
