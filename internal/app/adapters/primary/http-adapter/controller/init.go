package controller

import (
	"log/slog"

	api_service "github.com/mkorobovv/my-rest/internal/app/application/api-service"
)

type Controller struct {
	logger     *slog.Logger
	apiService *api_service.ApiService
}

func New(logger *slog.Logger, service *api_service.ApiService) *Controller {
	return &Controller{
		logger:     logger,
		apiService: service,
	}
}
