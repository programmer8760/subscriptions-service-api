package repository

import (
	"context"
	"database/sql"

	"github.com/programmer8760/subscriptions-service-api/internal/domain"
	"github.com/programmer8760/subscriptions-service-api/internal/dto"
)

type SubscriptionsRepository interface {
	Create(ctx context.Context, sub *domain.Subscription) error
	GetByID(ctx context.Context, id uint) (domain.Subscription, error)
	Update(ctx context.Context, req dto.UpdateSubscriptionDTO) (domain.Subscription, error)
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, req dto.GetAllSubscriptionsDTO) ([]domain.Subscription, error)
	GetTotalPrice(ctx context.Context, req dto.GetTotalPriceDTO) (int, error)
}

type PostgresSubscriptionsRepository struct {
	db *sql.DB
}

func NewPostgresSubscriptionsRepository(db *sql.DB) SubscriptionsRepository {
	return &PostgresSubscriptionsRepository{db: db}
}
