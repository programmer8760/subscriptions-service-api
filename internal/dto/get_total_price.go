package dto

import (
	"github.com/google/uuid"
	"github.com/prajkin/em-test-task/internal/domain"
)

type GetTotalPriceDTO struct {
	From   domain.Date
	To     domain.Date
	Name   *string
	UserID *uuid.UUID
}
