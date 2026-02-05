package service

import (
	"context"
	"strings"

	"github.com/prajkin/em-test-task/internal/domain"
	"github.com/prajkin/em-test-task/internal/dto"
)

func (s *SubscriptionsService) UpdateSubscription(ctx context.Context, req dto.UpdateSubscriptionDTO) error {
	if req.ID == 0 {
		return domain.ErrInvalidID
	}
	if req.Name == nil && req.Price == nil && req.UserID == nil && req.StartDate == nil && req.EndDate == nil {
		return domain.ErrNoChanges
	}
	if req.Name != nil {
		if *req.Name = strings.TrimSpace(*req.Name); *req.Name == "" {
			return domain.ErrInvalidName
		}
	}
	if req.Price != nil {
		if *req.Price <= 0 {
			return domain.ErrInvalidPrice
		}
	}
	if req.StartDate != nil {
		if (*req.StartDate).Time.IsZero() {
			return domain.ErrInvalidStartDate
		}
	}
	if req.EndDate != nil {
		if (*req.EndDate).Time.IsZero() {
			return domain.ErrInvalidEndDate
		}
	}

	if err := s.repo.Update(ctx, req); err != nil {
		return err
	}

	return nil
}
