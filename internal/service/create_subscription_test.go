package service

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/programmer8760/subscriptions-service-api/internal/domain"
)

func TestServiceCreateSubscription(t *testing.T) {
	tests := []struct {
		sub     domain.Subscription
		wantErr error
	}{
		{
			domain.Subscription{
				Name:      "",
				Price:     100,
				StartDate: domain.Date{time.Now()},
				UserID:    uuid.UUID{},
			}, domain.BadRequest{Err: domain.ErrInvalidName},
		}, {
			domain.Subscription{
				Name:      "test1",
				Price:     100,
				StartDate: domain.Date{time.Now()},
				UserID:    uuid.UUID{},
			}, nil,
		}, {
			domain.Subscription{
				Name:      "test2",
				Price:     0,
				StartDate: domain.Date{time.Now()},
				UserID:    uuid.UUID{},
			}, domain.BadRequest{Err: domain.ErrInvalidPrice},
		}, {
			domain.Subscription{
				Name:      "test3",
				Price:     -100,
				StartDate: domain.Date{time.Now()},
				UserID:    uuid.UUID{},
			}, domain.BadRequest{Err: domain.ErrInvalidPrice},
		}, {
			domain.Subscription{
				Name:      "test4",
				Price:     100,
				StartDate: domain.Date{time.Time{}},
				UserID:    uuid.UUID{},
			}, domain.BadRequest{Err: domain.ErrInvalidStartDate},
		}, {
			domain.Subscription{
				Name:      "test5",
				Price:     100,
				StartDate: domain.Date{time.Now()},
				UserID:    uuid.UUID{},
				EndDate:   &domain.Date{time.Now().Add(time.Hour * 24 * 31)},
			}, nil,
		}, {
			domain.Subscription{
				Name:      "test6",
				Price:     100,
				StartDate: domain.Date{time.Now()},
				UserID:    uuid.UUID{},
				EndDate:   &domain.Date{time.Time{}},
			}, domain.BadRequest{Err: domain.ErrInvalidEndDate},
		}, {
			domain.Subscription{
				Name:      "test7",
				Price:     100,
				StartDate: domain.Date{time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)},
				UserID:    uuid.UUID{},
				EndDate:   &domain.Date{time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)},
			}, domain.BadRequest{Err: domain.ErrEndBeforeStart},
		}, {
			domain.Subscription{
				Name:      "test8",
				Price:     100,
				StartDate: domain.Date{time.Now()},
				UserID:    uuid.UUID{},
				EndDate:   &domain.Date{time.Now().Add(-1 * time.Hour * 24 * 31)},
			}, domain.BadRequest{Err: domain.ErrEndBeforeStart},
		},
	}

	repo := &mockRepository{}
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	svc := NewSubscriptionsService(repo, log)

	for _, tt := range tests {
		t.Run(tt.sub.Name, func(t *testing.T) {
			_, err := svc.CreateSubscription(context.Background(), tt.sub)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
