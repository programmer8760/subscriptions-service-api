package service

import (
	"context"

	"github.com/prajkin/em-test-task/internal/domain"
)

func (s *SubscriptionsService) GetAllSubscriptions(ctx context.Context) ([]domain.Subscription, error) {
	subs, err := s.repo.List(ctx)

	return subs, err
}
