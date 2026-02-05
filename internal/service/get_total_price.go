package service

import (
	"context"
	"strings"

	"github.com/prajkin/em-test-task/internal/domain"
	"github.com/prajkin/em-test-task/internal/dto"
)

func (s *SubscriptionsService) GetTotalPrice(ctx context.Context, req dto.GetTotalPriceDTO) (int, error) {
	if req.From.Time.IsZero() {
		return 0, domain.ErrInvalidFromDate
	}
	if req.To.Time.IsZero() {
		return 0, domain.ErrInvalidToDate
	}
	if req.Name != nil {
		if *req.Name = strings.TrimSpace(*req.Name); *req.Name == "" {
			return 0, domain.ErrInvalidName
		}
	}

	total, err := s.repo.GetTotalPrice(ctx, req)
	if err != nil {
		return 0, err
	}

	return total, nil
}
