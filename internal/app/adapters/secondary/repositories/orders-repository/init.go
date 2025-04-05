package orders_repository

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type OrdersRepository struct {
	logger *slog.Logger
	DB     *sqlx.DB
}

func New(l *slog.Logger, db *sqlx.DB) *OrdersRepository {
	return &OrdersRepository{
		logger: l,
		DB:     db,
	}
}
