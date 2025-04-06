package orders_producer

import (
	"context"
	"encoding/json"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"github.com/mkorobovv/my-rest/internal/app/domain/order"
)

func (p *OrdersProducer) Produce(ctx context.Context, order order.Order) error {
	payload, err := json.Marshal(order)
	if err != nil {
		return err
	}

	e := cloudevents.NewEvent()

	e.SetTime(time.Now())
	e.SetType(p.config.Topic)
	e.SetSource("orders_generation_service")
	e.SetID(uuid.NewString())

	err = e.SetData(cloudevents.ApplicationJSON, payload)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	record := eventToRecord(e)

	err = p.client.ProduceSync(ctx, record).FirstErr()
	if err != nil {
		return err
	}

	return nil
}
