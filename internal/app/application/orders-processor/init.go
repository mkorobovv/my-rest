package orders_processor

import (
	"context"
	"log/slog"

	orders_generation_service "github.com/mkorobovv/my-rest/internal/app/application/orders-generation-service"
	"github.com/mkorobovv/my-rest/internal/app/domain/order"
)

type OrdersProcessor struct {
	logger                  *slog.Logger
	ordersGenerationService *orders_generation_service.OrdersGenerationService
	ordersProducer          ordersProducer
}

type ordersProducer interface {
	Produce(ctx context.Context, order order.Order) error
}

func New(l *slog.Logger, ordersProducer ordersProducer) *OrdersProcessor {
	return &OrdersProcessor{
		logger:                  l,
		ordersGenerationService: orders_generation_service.New(l),
		ordersProducer:          ordersProducer,
	}
}
