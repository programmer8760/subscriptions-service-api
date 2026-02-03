package repository

import (
	"context"
	"database/sql"

	"github.com/prajkin/em-test-task/internal/domain"
)

func (r *SubscriptionsRepository) List(ctx context.Context) ([]domain.Subscription, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT * FROM subscriptions",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []domain.Subscription
	for rows.Next() {
		var (
			sub     domain.Subscription
			endDate sql.NullTime
		)
		err := rows.Scan(&sub.ID, &sub.Name, &sub.Price, &sub.UserID, &sub.StartDate, &endDate)
		if err != nil {
			return nil, err
		}

		if endDate.Valid {
			sub.EndDate = &endDate.Time
		}
		subs = append(subs, sub)
	}

	return subs, nil
}
