package dto

import (
	"github.com/google/uuid"
	"github.com/programmer8760/subscriptions-service-api/internal/domain"
)

type GetTotalPriceDTO struct {
	From   domain.Date
	To     domain.Date
	Name   *string
	UserID *uuid.UUID
}
