package service

import (
	"context"

	"github.com/prajkin/em-test-task/internal/domain"
)

func (s *SubscriptionsService) DeleteSubscription(ctx context.Context, id uint) error {
	if id == 0 {
		return domain.ErrInvalidID
	}

	err := s.repo.Delete(ctx, id)

	return err
}
