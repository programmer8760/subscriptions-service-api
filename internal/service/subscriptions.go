package service

import (
	"context"
	"log/slog"

	"github.com/prajkin/em-test-task/internal/domain"
	"github.com/prajkin/em-test-task/internal/dto"
	"github.com/prajkin/em-test-task/internal/repository"
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
