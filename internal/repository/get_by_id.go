package repository

import (
	"context"

	"github.com/prajkin/em-test-task/internal/domain"
)

func (r *SubscriptionsRepository) GetByID(ctx context.Context, id uint) (domain.Subscription, error) {
	var sub domain.Subscription
	err := r.db.QueryRowContext(
		ctx,
		"SELECT * FROM subscriptions WHERE id = $1",
		id,
	).Scan(&sub.ID, &sub.Name, &sub.Price, &sub.UserID, &sub.StartDate, &sub.EndDate)
	return sub, err
}
