package orders_consumer

import (
	api_service "github.com/mkorobovv/my-rest/internal/app/application/api-service"
	"log/slog"

	"github.com/mkorobovv/my-rest/internal/app/adapters/primary/orders-consumer/handler"
	"github.com/mkorobovv/my-rest/internal/app/infrastructure/kafka"
	"github.com/twmb/franz-go/pkg/kgo"
)

type OrdersConsumer struct {
	logger  *slog.Logger
	client  *kgo.Client
	config  kafka.ConsumerConfig
	handler *handler.OrdersHandler
}

func New(l *slog.Logger, client *kgo.Client, cfg kafka.ConsumerConfig, apiService *api_service.ApiService) *OrdersConsumer {
	_handler := handler.New(l, apiService)

	client.AddConsumeTopics(cfg.Topic)

	cfg.BatchSize = 10

	return &OrdersConsumer{
		logger:  l,
		client:  client,
		config:  cfg,
		handler: _handler,
	}
}
