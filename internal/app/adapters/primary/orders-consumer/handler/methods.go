package handler

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func (h *OrdersHandler) Process(ctx context.Context, event cloudevents.Event) error {
	order, err := eventToOrder(event)
	if err != nil {
		return err
	}

	err = h.apiService.Create(ctx, order)
	if err != nil {
		return err
	}

	return nil
}
