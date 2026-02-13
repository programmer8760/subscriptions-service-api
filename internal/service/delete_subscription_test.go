package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"testing"

	"github.com/programmer8760/subscriptions-service-api/internal/domain"
)

func TestServiceDeleteSubscription(t *testing.T) {
	tests := []struct {
		id      uint
		wantErr error
	}{
		{0, domain.ErrInvalidID},
		{10, nil},
	}
	repo := &mockRepository{}
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	svc := NewSubscriptionsService(repo, log)

	for _, tt := range tests {
		name := fmt.Sprintf("id %d", tt.id)
		t.Run(name, func(t *testing.T) {
			err := svc.DeleteSubscription(context.Background(), tt.id)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("err = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
