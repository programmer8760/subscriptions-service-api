package dto

import (
	"time"

	"github.com/google/uuid"
)

type UpdateSubscriptionDTO struct {
	ID        uint
	Name      *string
	Price     *int
	UserID    *uuid.UUID
	StartDate *time.Time
	EndDate   *time.Time
}
