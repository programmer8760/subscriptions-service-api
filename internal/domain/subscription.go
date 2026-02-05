package domain

import (
	"github.com/google/uuid"
)

type Subscription struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	UserID    uuid.UUID `json:"user_id"`
	StartDate Date      `json:"start_date"`
	EndDate   *Date     `json:"end_date,omitempty"`
}
