package service

import "github.com/prajkin/em-test-task/internal/repository"

type SubscriptionsService struct {
	repo *repository.SubscriptionsRepository
}

func NewSubscriptionsService(repo *repository.SubscriptionsRepository) *SubscriptionsService {
	return &SubscriptionsService{repo: repo}
}
