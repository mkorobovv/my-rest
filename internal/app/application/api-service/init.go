package api_service

import (
	"context"
	"github.com/kofalt/go-memoize"
	"log/slog"
	"time"

	"github.com/mkorobovv/my-rest/internal/app/domain/order"
)

type ApiService struct {
	logger           *slog.Logger
	ordersRepository ordersRepository
	cache            *memoize.Memoizer
}

type ordersRepository interface {
	Create(ctx context.Context, order order.Order) (err error)
	Get(ctx context.Context, trackNumber string) (order order.Order, err error)
	Update(ctx context.Context, order order.Order) (err error)
}

func New(l *slog.Logger, ordersRepository ordersRepository) *ApiService {
	return &ApiService{
		logger:           l,
		ordersRepository: ordersRepository,
		cache:            memoize.NewMemoizer(2*time.Hour, 4*time.Hour),
	}
}
