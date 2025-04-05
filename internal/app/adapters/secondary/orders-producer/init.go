package orders_producer

import (
	"log/slog"

	"github.com/mkorobovv/my-rest/internal/app/infrastructure/kafka"
	"github.com/twmb/franz-go/pkg/kgo"
)

type OrdersProducer struct {
	logger *slog.Logger
	client *kgo.Client
	config kafka.ProducerConfig
}

func New(l *slog.Logger, cfg kafka.ProducerConfig, client *kgo.Client) *OrdersProducer {
	return &OrdersProducer{
		logger: l,
		client: client,
		config: cfg,
	}
}
