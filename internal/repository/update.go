package repository

import (
	"context"
	"strconv"
	"strings"

	"github.com/prajkin/em-test-task/internal/domain"
	"github.com/prajkin/em-test-task/internal/dto"
)

func (r *SubscriptionsRepository) Update(ctx context.Context, req dto.UpdateSubscriptionDTO) error {
	var builder strings.Builder
	var args []any

	builder.WriteString("UPDATE subscriptions SET ")
	if req.Name != nil {
		builder.WriteString("name = $")
		builder.WriteString(strconv.Itoa(len(args) + 1))
		builder.WriteRune(' ')
		args = append(args, req.Name)
	}
	if req.Price != nil {
		builder.WriteString("price = $")
		builder.WriteString(strconv.Itoa(len(args) + 1))
		builder.WriteRune(' ')
		args = append(args, req.Price)
	}
	if req.UserID != nil {
		builder.WriteString("user_id = $")
		builder.WriteString(strconv.Itoa(len(args) + 1))
		builder.WriteRune(' ')
		args = append(args, req.UserID)
	}
	if req.StartDate != nil {
		builder.WriteString("start_date = $")
		builder.WriteString(strconv.Itoa(len(args) + 1))
		builder.WriteRune(' ')
		args = append(args, req.StartDate.Time)
	}
	if req.EndDate != nil {
		builder.WriteString("end_date = $")
		builder.WriteString(strconv.Itoa(len(args) + 1))
		builder.WriteRune(' ')
		args = append(args, req.EndDate.Time)
	}
	builder.WriteString("WHERE id = $")
	builder.WriteString(strconv.Itoa(len(args) + 1))
	args = append(args, req.ID)

	res, err := r.db.ExecContext(
		ctx,
		builder.String(),
		args...,
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
