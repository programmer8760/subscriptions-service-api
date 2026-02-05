package repository

import (
	"context"

	"github.com/prajkin/em-test-task/internal/domain"
	"github.com/prajkin/em-test-task/internal/dto"
)

func (r *SubscriptionsRepository) Update(ctx context.Context, req dto.UpdateSubscriptionDTO) error {
	res, err := r.db.ExecContext(
		ctx,
		`UPDATE subscriptions SET
		name=COALESCE($1, name),
		price=COALESCE($2, price),
		user_id=COALESCE($3, user_id),
		start_date=COALESCE($4, start_date),
		end_date=COALESCE($5, end_date)
		WHERE id=$6`,
		req.Name,
		req.Price,
		req.UserID,
		req.StartDate,
		req.EndDate,
		req.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrSubscriptionNotFound
	}

	return nil
}
