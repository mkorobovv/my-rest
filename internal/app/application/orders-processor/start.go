package orders_processor

import (
	"context"
	"time"
)

func (svc *OrdersProcessor) Start(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return nil
	default:
	}

	const sleepTime = 5 * time.Second

	for {
		order := svc.ordersGenerationService.Generate()

		err := svc.ordersProducer.Produce(ctx, order)
		if err != nil {
			return err
		}

		select {
		case <-time.After(sleepTime):
		case <-ctx.Done():
			return nil
		}
	}
}
