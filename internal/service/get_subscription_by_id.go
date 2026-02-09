package service

import (
	"context"

	"github.com/programmer8760/subscriptions-service-api/internal/domain"
)

func (s *subscriptionsService) GetSubscriptionByID(ctx context.Context, id uint) (domain.Subscription, error) {
	if id == 0 {
		return domain.Subscription{}, domain.ErrInvalidID
	}

	sub, err := s.repo.GetByID(ctx, id)

	return sub, err
}
