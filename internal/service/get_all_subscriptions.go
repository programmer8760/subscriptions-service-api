package service

import (
	"context"

	"github.com/programmer8760/subscriptions-service-api/internal/domain"
	"github.com/programmer8760/subscriptions-service-api/internal/dto"
)

func (s *subscriptionsService) GetAllSubscriptions(ctx context.Context, req dto.GetAllSubscriptionsDTO) ([]domain.Subscription, error) {
	subs, err := s.repo.List(ctx, req)
	if err != nil {
		return nil, err
	}

	return subs, nil
}
