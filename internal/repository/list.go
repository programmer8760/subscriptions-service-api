package repository

import (
	"context"
	"database/sql"

	"github.com/programmer8760/subscriptions-service-api/internal/domain"
	"github.com/programmer8760/subscriptions-service-api/internal/dto"
)

func (r *PostgresSubscriptionsRepository) List(ctx context.Context, req dto.GetAllSubscriptionsDTO) ([]domain.Subscription, error) {
	limit := domain.DefaultPageSize
	if req.PageSize != nil {
		limit = *req.PageSize
	}
	var offset uint
	if req.Page != nil {
		offset = (*req.Page - 1) * limit
	}

	rows, err := r.db.QueryContext(
		ctx,
		"SELECT * FROM subscriptions LIMIT $1 OFFSET $2",
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []domain.Subscription
	for rows.Next() {
		var (
			sub       domain.Subscription
			startDate sql.NullTime
			endDate   sql.NullTime
		)
		err := rows.Scan(&sub.ID, &sub.Name, &sub.Price, &sub.UserID, &startDate, &endDate)
		if err != nil {
			return nil, err
		}

		sub.StartDate = domain.NewDate(startDate.Time)
		if endDate.Valid {
			t := domain.NewDate(endDate.Time)
			sub.EndDate = &t
		}
		subs = append(subs, sub)
	}

	return subs, nil
}
