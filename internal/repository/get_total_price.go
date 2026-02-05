package repository

import (
	"context"

	"github.com/prajkin/em-test-task/internal/dto"
)

func (r *SubscriptionsRepository) GetTotalPrice(ctx context.Context, req dto.GetTotalPriceDTO) (int, error) {
	var total int

	err := r.db.QueryRowContext(
		ctx,
		`SELECT COALESCE(SUM(s.price), 0)
		FROM subscriptions s
		JOIN generate_series(
			$1::date,
			($2::date - interval '1 month'),
			interval '1 month'
		) m(month_start)
		ON s.start_date <= m.month_start
		AND (s.end_date IS NULL OR s.end_date > m.month_start)
		WHERE ($3::uuid IS NULL OR s.user_id = $3)
			AND ($4::text IS NULL OR s.name = $4)`,
		req.From.Time,
		req.To.Time,
		req.UserID,
		req.Name,
	).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
