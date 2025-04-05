package orders_generation_service

import (
	"log/slog"
)

type OrdersGenerationService struct {
	logger *slog.Logger
}

func New(l *slog.Logger) *OrdersGenerationService {
	return &OrdersGenerationService{
		logger: l,
	}
}
