package service

import (
	"log/slog"

	"github.com/prajkin/em-test-task/internal/repository"
)

type SubscriptionsService struct {
	repo   *repository.SubscriptionsRepository
	logger *slog.Logger
}

func NewSubscriptionsService(repo *repository.SubscriptionsRepository, log *slog.Logger) *SubscriptionsService {
	return &SubscriptionsService{
		repo:   repo,
		logger: log,
	}
}
