package service

import (
	"context"
	"log/slog"

	"github.com/programmer8760/subscriptions-service-api/internal/domain"
	"github.com/programmer8760/subscriptions-service-api/internal/dto"
	"github.com/programmer8760/subscriptions-service-api/internal/repository"
)

type SubscriptionsService interface {
	CreateSubscription(ctx context.Context, sub domain.Subscription) (domain.Subscription, error)
	DeleteSubscription(ctx context.Context, id uint) error
	GetAllSubscriptions(ctx context.Context) ([]domain.Subscription, error)
	GetSubscriptionByID(ctx context.Context, id uint) (domain.Subscription, error)
	GetTotalPrice(ctx context.Context, req dto.GetTotalPriceDTO) (int, error)
	UpdateSubscription(ctx context.Context, req dto.UpdateSubscriptionDTO) (domain.Subscription, error)
}

type subscriptionsService struct {
	repo   repository.SubscriptionsRepository
	logger *slog.Logger
}

func NewSubscriptionsService(repo repository.SubscriptionsRepository, log *slog.Logger) SubscriptionsService {
	return &subscriptionsService{
		repo:   repo,
		logger: log,
	}
}
