package handler

import (
	"log/slog"

	api_service "github.com/mkorobovv/my-rest/internal/app/application/api-service"
)

type OrdersHandler struct {
	logger     *slog.Logger
	apiService *api_service.ApiService
}

func New(logger *slog.Logger, apiService *api_service.ApiService) *OrdersHandler {
	return &OrdersHandler{
		logger:     logger,
		apiService: apiService,
	}
}
