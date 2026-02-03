package repository

import (
	"context"
	"database/sql"

	"github.com/prajkin/em-test-task/internal/domain"
)

func (r *SubscriptionsRepository) GetByID(ctx context.Context, id uint) (domain.Subscription, error) {
	var (
		sub     domain.Subscription
		endDate sql.NullTime
	)
	err := r.db.QueryRowContext(
		ctx,
		"SELECT * FROM subscriptions WHERE id = $1",
		id,
	).Scan(&sub.ID, &sub.Name, &sub.Price, &sub.UserID, &sub.StartDate, &endDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Subscription{}, domain.ErrSubscriptionNotFound
		}
		return domain.Subscription{}, err
	}

	if endDate.Valid {
		sub.EndDate = &endDate.Time
	}

	return sub, nil
}
