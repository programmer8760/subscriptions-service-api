package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/programmer8760/subscriptions-service-api/internal/domain"
	"github.com/programmer8760/subscriptions-service-api/internal/dto"
)

func TestServiceGetTotalPrice(t *testing.T) {
	tests := []struct {
		dto     dto.GetTotalPriceDTO
		wantErr error
		id      uint
	}{
		{
			dto.GetTotalPriceDTO{
				From: domain.Date{time.Now()},
				To:   domain.Date{time.Time{}},
			}, domain.BadRequest{Err: domain.ErrInvalidToDate}, 0,
		}, {
			dto.GetTotalPriceDTO{
				From: domain.Date{time.Time{}},
				To:   domain.Date{time.Now()},
			}, domain.BadRequest{Err: domain.ErrInvalidFromDate}, 1,
		}, {
			dto.GetTotalPriceDTO{
				From: domain.Date{time.Now()},
				To:   domain.Date{time.Now()},
				Name: ptr(""),
			}, domain.BadRequest{Err: domain.ErrInvalidName}, 2,
		}, {
			dto.GetTotalPriceDTO{
				From: domain.Date{time.Now()},
				To:   domain.Date{time.Now()},
				Name: ptr("test"),
			}, nil, 3,
		},
	}

	repo := &mockRepository{}
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	svc := NewSubscriptionsService(repo, log)

	for _, tt := range tests {
		name := fmt.Sprintf("%d", tt.id)
		t.Run(name, func(t *testing.T) {
			_, err := svc.GetTotalPrice(context.Background(), tt.dto)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
