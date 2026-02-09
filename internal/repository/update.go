package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/prajkin/em-test-task/internal/domain"
	"github.com/prajkin/em-test-task/internal/dto"
)

func (r *SubscriptionsRepository) Update(ctx context.Context, req dto.UpdateSubscriptionDTO) (domain.Subscription, error) {
	var sub domain.Subscription
	err := r.db.QueryRowContext(
		ctx,
		`UPDATE subscriptions SET
		name=COALESCE($1, name),
		price=COALESCE($2, price),
		user_id=COALESCE($3, user_id),
		start_date=COALESCE($4, start_date),
		end_date=COALESCE($5, end_date)
		WHERE id=$6
		RETURNING *`,
		req.Name,
		req.Price,
		req.UserID,
		req.StartDate,
		req.EndDate,
		req.ID,
	).Scan(&sub)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Subscription{}, domain.ErrSubscriptionNotFound
		}
		return domain.Subscription{}, err
	}

	return sub, nil
}
