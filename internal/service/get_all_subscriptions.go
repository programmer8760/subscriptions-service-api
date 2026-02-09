package service

import (
	"context"

	"github.com/programmer8760/subscriptions-service-api/internal/domain"
)

func (s *subscriptionsService) GetAllSubscriptions(ctx context.Context) ([]domain.Subscription, error) {
	subs, err := s.repo.List(ctx)

	return subs, err
}
