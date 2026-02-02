package repository

import (
	"context"

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
		var sub domain.Subscription
		err := rows.Scan(&sub.ID, &sub.Name, &sub.Price, &sub.UserID, &sub.StartDate, &sub.EndDate)
		if err != nil {
			return nil, err
		}
		subs = append(subs, sub)
	}

	return subs, nil
}
