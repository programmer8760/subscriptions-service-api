package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/programmer8760/subscriptions-service-api/internal/domain"
	"github.com/programmer8760/subscriptions-service-api/internal/dto"
)

func ptr[T any](v T) *T { return &v }

func TestServiceUpdateSubscription(t *testing.T) {
	tests := []struct {
		dto     dto.UpdateSubscriptionDTO
		wantErr error
	}{
		{
			dto.UpdateSubscriptionDTO{
				ID: 0,
			}, domain.BadRequest{Err: domain.ErrInvalidID},
		}, {
			dto.UpdateSubscriptionDTO{
				ID:   1,
				Name: ptr(""),
			}, domain.BadRequest{Err: domain.ErrInvalidName},
		}, {
			dto.UpdateSubscriptionDTO{
				ID:    2,
				Price: ptr(0),
			}, domain.BadRequest{Err: domain.ErrInvalidPrice},
		}, {
			dto.UpdateSubscriptionDTO{
				ID:        3,
				StartDate: ptr(domain.Date{time.Time{}}),
			}, domain.BadRequest{Err: domain.ErrInvalidStartDate},
		}, {
			dto.UpdateSubscriptionDTO{
				ID:      4,
				EndDate: ptr(domain.Date{time.Time{}}),
			}, domain.BadRequest{Err: domain.ErrInvalidEndDate},
		}, {
			dto.UpdateSubscriptionDTO{
				ID: 5,
			}, domain.BadRequest{Err: domain.ErrNoChanges},
		}, {
			dto.UpdateSubscriptionDTO{
				ID:        6,
				Name:      ptr("test6"),
				Price:     ptr(100),
				StartDate: ptr(domain.Date{time.Now()}),
				UserID:    ptr(uuid.UUID{}),
				EndDate:   ptr(domain.Date{time.Now()}),
			}, nil,
		},
	}

	repo := &mockRepository{}
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	svc := NewSubscriptionsService(repo, log)

	for _, tt := range tests {
		name := fmt.Sprintf("%d", tt.dto.ID)
		t.Run(name, func(t *testing.T) {
			_, err := svc.UpdateSubscription(context.Background(), tt.dto)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
