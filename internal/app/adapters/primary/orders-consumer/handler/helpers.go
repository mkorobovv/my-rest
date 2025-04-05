package handler

import (
	"encoding/json"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/mkorobovv/my-rest/internal/app/domain/order"
)

func eventToOrder(event cloudevents.Event) (_order order.Order, err error) {
	err = json.Unmarshal(event.Data(), &_order)
	if err != nil {
		return order.Order{}, err
	}

	return _order, nil
}
