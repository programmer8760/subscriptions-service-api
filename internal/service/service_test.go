package service

import (
	"context"

	"github.com/programmer8760/subscriptions-service-api/internal/domain"
	"github.com/programmer8760/subscriptions-service-api/internal/dto"
)

type mockRepository struct{}

func (m *mockRepository) Create(ctx context.Context, sub *domain.Subscription) error {
	return nil
}

func (m *mockRepository) GetByID(ctx context.Context, id uint) (domain.Subscription, error) {
	return domain.Subscription{}, nil
}

func (m *mockRepository) Update(ctx context.Context, req dto.UpdateSubscriptionDTO) (domain.Subscription, error) {
	return domain.Subscription{}, nil
}

func (m *mockRepository) Delete(ctx context.Context, id uint) error {
	return nil
}

func (m *mockRepository) List(ctx context.Context, req dto.GetAllSubscriptionsDTO) ([]domain.Subscription, error) {
	return nil, nil
}

func (m *mockRepository) GetTotalPrice(ctx context.Context, req dto.GetTotalPriceDTO) (int, error) {
	return 0, nil
}
