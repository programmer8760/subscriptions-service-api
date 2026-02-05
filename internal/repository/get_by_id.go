package repository

import (
	"context"
	"database/sql"

	"github.com/prajkin/em-test-task/internal/domain"
)

func (r *SubscriptionsRepository) GetByID(ctx context.Context, id uint) (domain.Subscription, error) {
	var (
		sub       domain.Subscription
		startDate sql.NullTime
		endDate   sql.NullTime
	)
	err := r.db.QueryRowContext(
		ctx,
		"SELECT * FROM subscriptions WHERE id = $1",
		id,
	).Scan(&sub.ID, &sub.Name, &sub.Price, &sub.UserID, &startDate, &endDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Subscription{}, domain.ErrSubscriptionNotFound
		}
		return domain.Subscription{}, err
	}

	sub.StartDate = domain.NewDate(startDate.Time)
	if endDate.Valid {
		t := domain.NewDate(endDate.Time)
		sub.EndDate = &t
	}

	return sub, nil
}
