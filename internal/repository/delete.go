package repository

import (
	"context"

	"github.com/prajkin/em-test-task/internal/domain"
)

func (r *SubscriptionsRepository) Delete(ctx context.Context, id uint) error {
	res, err := r.db.ExecContext(
		ctx,
		"DELETE FROM subscriptions WHERE id = $1",
		id,
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
