package repository

import (
	"context"
	"time"

	"github.com/prajkin/em-test-task/internal/domain"
)

func (r *SubscriptionsRepository) Create(ctx context.Context, sub *domain.Subscription) error {
	var endDate *time.Time
	if sub.EndDate != nil {
		endDate = &sub.EndDate.Time
	}
	err := r.db.QueryRowContext(
		ctx,
		"INSERT INTO subscriptions (name, price, user_id, start_date, end_date) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		sub.Name,
		sub.Price,
		sub.UserID,
		sub.StartDate.Time,
		endDate,
	).Scan(&sub.ID)
	return err
}
